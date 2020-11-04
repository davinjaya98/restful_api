package router

import (
	"github.com/gorilla/mux"

	_ "github.com/davinjaya98/restful_api/docs" // docs is generated by Swag CLI, you have to import it.

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/davinjaya98/restful_api/internal/contacts"
)

// Main function
func InitRouter() *mux.Router{
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	// Get all
	r.HandleFunc("/contacts", contacts.GetAll).Methods("GET")
	// Get
	r.HandleFunc("/contacts/{name}", contacts.Get).Methods("GET")
	// Create
	r.HandleFunc("/contacts", contacts.Create).Methods("POST")
	// Update
	r.HandleFunc("/contacts/{name}", contacts.Update).Methods("PUT")
	// Delete
	r.HandleFunc("/contacts/{name}", contacts.Delete).Methods("DELETE")

	// Swagger
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return r;
}