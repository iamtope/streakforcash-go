package auth 

import (
	"fmt"
	"os"
	"net/http"
	"strings"
	"streakforcash-api-go-version/app/models"
	jwt "github.com/dgrijalva/jwt-go"
	u "streakforcash-api-go-version/app/utils"


)

var jwtAuthentication = func (next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
		noAuth := []string{ "/api/user/login"}
		requestPath := r.URL.Path
		fmt.Print(requestPath)

		for _,  v := range noAuth {
			if v == requestPath {
				next.ServeHTTP(w, r)
				return // comment this return out, it seems to be unneccessary
			}

		}
		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")

		// if token header is missing
		if tokenHeader == "" {
			response = u.Message(false, "missing token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-type", "application/json")
			u.Respond(w, response)
			return
		}
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.Username) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

