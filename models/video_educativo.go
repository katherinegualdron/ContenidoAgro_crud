package routes

import (
	"CONTENIDO/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasVideoEducativo(router *mux.Router) {
	router.HandleFunc("/video_educativo", controllers.ObtenerVideosEducativos).Methods("GET")
	router.HandleFunc("/video_educativo/{id}", controllers.ObtenerVideoEducativoPorID).Methods("GET")
	router.HandleFunc("/video_educativo", controllers.CrearVideoEducativo).Methods("POST")
	router.HandleFunc("/video_educativo/{id}", controllers.ActualizarVideoEducativo).Methods("PUT")
	router.HandleFunc("/video_educativo/{id}", controllers.EliminarVideoEducativo).Methods("DELETE")
}
