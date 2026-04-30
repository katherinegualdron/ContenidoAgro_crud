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
	routes.RegistrarRutas(router)

	log.Println("Servidor CONTENIDO escuchando en :8092")
	log.Fatal(http.ListenAndServe(":8092", router))
}
