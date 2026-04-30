package routes

import (
	"CONTENIDO/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasVideoEducativo(router *mux.Router) {
	
router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
