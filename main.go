package main

import (
	"log"
	"net/http"

	"API_GO_CRUD/config"
	"API_GO_CRUD/routes"

	"github.com/gorilla/mux" // Libreria para crear rutas
)

// middieware CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB() // Conectar a la base de datos

	r := mux.NewRouter()
	// Registro de rutas
	routes.RegisterCredentialRoutes(r)
	routes.RegisterUserRoutes(r)

	log.Println("Servidor corriendo en el puerto 8083")

	http.ListenAndServe(":8083", enableCors(r)) // Iniciar servidor con CORS
}