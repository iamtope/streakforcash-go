package models

import (
	"github.com/dgrijalva/jwt-go"
	u "streakforcash-api-go-version/app/utils"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"os"
	"log"
	"fmt"
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
   fmt.Println("i actually got here, means everythung is fine")
   uniqueID, _ := uuid.NewUUID()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	createUser := `
	INSERT INTO 
	user_info(id, role, username, email, password )
	VALUES
	(?,?,?,?,?)`
	// newUser := &CreateUser{Username: user.Username, Email: user.Email, Password: user.Password, Role: "basic", ID: uniqueID}
	stmt, err := dbx.Prepare(createUser)
	if err != nil {
		log.Println("Error preparing statement")
		panic(err)
	}

log.Println("Incoming user object: ", user)
	row, err := stmt.Exec(uniqueID, "basic", "iamtope", "iamtope@gmail.com", "$2a$10$ZGQa/tgQzTkD9QJQhXOBletUqTvDx6hFpj/vr6yy/DLC0M6i26v86")
	log.Println("Row is: ")
	// err := row.Scan(newUser)
	if err != nil {
		switch err {
        case sql.ErrNoRows:
            return u.Message(false, "not found" )
        default:
            return u.Message(false, "an error occured" )
		}
	}

	log.Println("Row is: ", row)
	//Create new JWT token for the newly registered User
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("secret")))
	user.Token = tokenString

	user.Password = "" //delete password


	resp := u.Message(true, "User created successfully")
	return resp
}