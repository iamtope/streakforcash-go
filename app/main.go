package main

import (
	"os"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"

)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}


fmt.Println(port)

err := http.ListenAndServe(":" + port, r)

if err != nil {
	fmt.Print(err)
}

}