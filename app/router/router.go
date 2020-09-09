package router

import(
	"github.com/gorilla/mux"
	"streakforcash-api-go-version/app/controllers"
)

// router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/signup", controllers.createUser ).Methods("POST", "OPTIONS")

	return router
}