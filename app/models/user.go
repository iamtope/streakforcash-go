package models

import (
	"github.com/dgrijalva/jwt-go"
	u "streakforcash-api-go-version/app/utils"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"os"
)


// jwt struct 

type Token struct {
	UserId int
	jwt.StandardClaims
}
// User structure
type User  struct {
	ID int `json: "id"`
	Email string `json: "email"`
	Password string `json: "password"`
	Role string `json: "password`
	Token string `json: "token"`
	Username string `json: "username"`
	CreatedAt  string`json: "createdAt"`
	UpdatedAt  string `json: "updatedAt"`
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

func (user *User) Create() (map[string] interface{}){
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	createUser := `
			INSERT INTO 
				user(id, role, username, email, created_at, updated_at, password )
			VALUES
				($1,$2,$3,$4,$5,$6,$7) 
			RETURNING 
				id;`
	
	row := db.QueryRow(createUser, user.ID, user.Role, user.Email, user.CreatedAt)
	err := row.Scan(&user.ID)
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