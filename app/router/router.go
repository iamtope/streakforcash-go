package router

import(
	"github.com/gorilla/mux"
	controllers "streakforcash-api-go-version/app/controllers"
)

// router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/signup", controllers.CreateUser ).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/login", controllers.LoginUser ).Methods("POST", "OPTIONS")

	return router
}