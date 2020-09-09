package controllers

import (
	"net/http"
	"streakforcash-api-go-version/app/models"
	"encoding/json"
	u "streakforcash-api-go-version/app/utils"
)

var createUser = func(w http.ResponseWriter, r *http.Request) {
	userResponse := r.Context().Value("user") . (int) //Grab the id of the user that send the request
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Response(w, u.Message(false, "Error while decoding request body"))
		return
	}

	user.ID = userResponse
	resp := user.Create()
	u.Response(w, resp)

}