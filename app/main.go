package main

import (
	"os"
	"net/http"
	"fmt"
	"streakforcash-api-go-version/app/router"

)

func main() {
	r := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}


err := http.ListenAndServe(":" + port, r)

if err != nil {
	fmt.Print(err)
}

}