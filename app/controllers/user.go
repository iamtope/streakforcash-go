package controllers

import (
	"net/http"
	"log"

	"streakforcash-api-go-version/app/models"
	"encoding/json"
	u "streakforcash-api-go-version/app/utils"
)
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	data := &models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		panic(err)
    }
	log.Println("data is", data)
	d := models.User{Email: data.Email, Username: data.Username, Password: data.Password}
	resp := d.Create()
	u.Response(w, resp)
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
 
	json.NewEncoder(w).Encode(data)

}