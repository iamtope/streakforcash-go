package controllers

import (
	"net/http"
	"log"

	"streakforcash-api-go-version/app/models"
	"encoding/json"
	u "streakforcash-api-go-version/app/utils"
)
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	// user := &models.User{}
	data := &models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		panic(err)
    }
	log.Println(models.User{}, "pointer location here")
	log.Println(data)
	d := models.User{Email: data.Email, Username: data.Username, Password: data.Password}
	resp := d.Create()
	u.Response(w, resp)
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
 
	json.NewEncoder(w).Encode(data)


	// err := json.NewDecoder(r.Body).Decode(user)
	// if err != nil {
	// 	u.Response(w, u.Message(false, "Error while decoding request body"))
	// 	return
	// }

	// user.ID = userResponse
	// resp := user.Create()
	// u.Response(w, resp)

}