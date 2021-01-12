package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)


// EncodeAuthToken signs authentication token
func EncodeToken(uid string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = uid
	claims["IssuedAt"] = time.Now().Unix()
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("H256"), claims)
	return token.SignedString([]byte(os.Getenv("secret")))
}