package utils

import (
	"encoding/json"
	"net/http"
	"errors"

	"golang.org/x/crypto/bcrypt"
)


func Message (status bool, message string) (map[string] interface{}){
	return map[string] interface{} { "message" : message, "status" : status }
}

// had to add this comment because of go-lint complaint ...
func Response (w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

// had to add this comment because of go-lint complaint ...
func CheckPasswordHash ( password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("Incorrect Password, try again")
	}
	return nil
}