package routes

import (
	"CONTENIDO/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasRespuestaForo(router *mux.Router) {
	router.HandleFunc("/respuesta_foro", controllers.ObtenerRespuestasForo).Methods("GET")
	router.HandleFunc("/respuesta_foro/{id}", controllers.ObtenerRespuestaForoPorID).Methods("GET")
	router.HandleFunc("/respuesta_foro", controllers.CrearRespuestaForo).Methods("POST")
	router.HandleFunc("/respuesta_foro/{id}", controllers.ActualizarRespuestaForo).Methods("PUT")
	router.HandleFunc("/respuesta_foro/{id}", controllers.EliminarRespuestaForo).Methods("DELETE")
}
