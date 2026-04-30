package routes

import (
	"CONTENIDO/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasTemaForo(router *mux.Router) {
	router.HandleFunc("/tema_foro", controllers.ObtenerTemasForo).Methods("GET")
	router.HandleFunc("/tema_foro/{id}", controllers.ObtenerTemaForoPorID).Methods("GET")
	router.HandleFunc("/tema_foro", controllers.CrearTemaForo).Methods("POST")
	router.HandleFunc("/tema_foro/{id}", controllers.ActualizarTemaForo).Methods("PUT")
	router.HandleFunc("/tema_foro/{id}", controllers.EliminarTemaForo).Methods("DELETE")
}
