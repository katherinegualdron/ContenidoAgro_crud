package routes

import (
	"API_GO_CRUD/controllers"

	"github.com/gorilla/mux" 
)

func RegisterCredentialRoutes(r *mux.Router){
	r.HandleFunc("/credentials", controllers.GetALLCredentials).Methods("GET")
	r.HandleFunc("/credentials/{id}", controllers.GetCredentialByID).Methods("GET")
	r.HandleFunc("/credentials", controllers.CreateCredential).Methods("POST")
	r.HandleFunc("/credentials/{id}", controllers.UpdateCredential).Methods("PUT")
	r.HandleFunc("/credentials/{id}", controllers.DeleteCredential).Methods("DELETE")
}