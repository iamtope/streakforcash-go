package models

import (
	"github.com/dgrijalva/jwt-go"
	u "streakforcash-api-go-version/app/utils"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"os"
	"log"
	"github.com/google/uuid"
)


// jwt struct 

type Token struct {
	UserId string
	jwt.StandardClaims
}
// User structure
type User  struct {
	ID string `json:"id"`
	Email string `bson:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	Token string `json:"token"`
	Username string `json:"username"`
}

// validation
func (user *User) Validate() (map[string] interface{}, bool){
	if !strings.Contains(user.Email, "@"){
		return u.Message(false, "Email format is invalid, please enter a valid email"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "length of password must me less than or greater than 6"), false
	}

	// Email must be unique
	CheckUniqueEmail := &User{}
	if CheckUniqueEmail.Email != "" {
		return u.Message(false, "A user is registered has registered with us using this email, please try another"), false
	} 
	return u.Message(false, "Requirement passed"), true
}

type CreateUser struct {
	ID string
	Username string
	Email string
	Password string
	Role string
}


func (user *User) Create() (map[string] interface{}){
	var dbx = ReturnInstance()
	if resp, ok := user.Validate(); !ok {
		return resp
	}
   uniqueID, _ := uuid.NewUUID()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	createUser := `
	INSERT INTO 
	user_info(id, role, username, email, password )
	VALUES ($1, $2, $3, $4, $5)`
	// newUser := &CreateUser{Username: user.Username, Email: user.Email, Password: user.Password, Role: "basic", ID: uniqueID}
	_, err := dbx.Exec(createUser, uniqueID, "basic", user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error preparing statement", err)
		panic(err)
	}
	if err != nil {
		switch err {
        case sql.ErrNoRows:
            return u.Message(false, "not found" )
        default:
            return u.Message(false, "an error occured" )
		}
	}

 	//Create new JWT token for the newly registered User
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("secret")))
	user.Token = tokenString

	user.Password = "" //delete password


	resp := u.Message(true, "User created successfully")
	return resp
}

// login user
func (user *User) Login() (map[string] interface{}){
	var dbx = ReturnInstance()
	findUser := `SELECT * FROM user_info WHERE email=$1`
	usr, _ := dbx.Exec(findUser, user.Email)
	log.Println("user is", usr)
	if usr == nil {
		return u.Message(false, "You don't seem to be a part of us yet! Why not sign up and let's roll")
	}
	err := u.CheckPasswordHash(user.Email, user.Password)
	if err != nil {
        return u.Message(false, "Login Failed, Please try again")
	}
	token, err := u.EncodeToken(user.ID)
	user.Token = token
	resp := u.Message(true, "User token decoded successfully")
	return resp
}