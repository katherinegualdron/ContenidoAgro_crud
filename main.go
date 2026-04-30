package main

import (
	"log"
	"net/http"

	"CONTENIDO/config"
	"CONTENIDO/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()
    routes.RegistrarRutasVideoEducativo(router)
	log.Println("Servidor CONTENIDO escuchando en :8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}
