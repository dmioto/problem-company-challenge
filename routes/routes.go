// Router for the app
package routes

import (
	"back/services"

	"github.com/gorilla/mux"
)

// API exposed endpoits are defined here
func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	// Get an array up to 50 users
	router.HandleFunc("/users", services.GetUsers).Methods("GET")
	// Get a specific user by id
	router.HandleFunc("/user/{id}", services.GetUser).Methods("GET")
	// Create a user
	router.HandleFunc("/user", services.CreateUser).Methods("POST")
	// Update a user
	router.HandleFunc("/user/{id}", services.UpdateUser).Methods("PUT")
	// Delete a user
	router.HandleFunc("/user/{id}", services.DeleteUser).Methods("DELETE")

	return router
}
