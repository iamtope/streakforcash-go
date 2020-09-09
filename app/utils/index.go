package utils

import (
	"encoding/json"
	"net/http"
)


func Message (status bool, message string) (map[string] interface{}){
	return map[string] interface{} { "message" : message, "status" : status }
}

// had to add this comment because of go-lint complaint ...
func Response (w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}