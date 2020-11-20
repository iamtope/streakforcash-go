package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"streakforcash-api-go-version/app/router"

)

func main() {
	r := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
log.Println("Server up and running on port: ", port)
err := http.ListenAndServe(":" + port, r)
if err != nil {
	fmt.Print(err)
}
}