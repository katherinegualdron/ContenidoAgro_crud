package controllers

import (
	"API_GO_CRUD/config"
	"API_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" //Librería para crear rutas
)

// Helper respuesta JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Get ALL

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	rows, err := config.DB.Query("SELECT id, name, email, age, activo FROM users")

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error": err.Error()})
		return
	}

	var list []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.Activo)
		list = append(list, u)
	}

	respondJSON(w, 200, list)
}

// Get by Id

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var u models.User

	err := config.DB.QueryRow(
		"SELECT id, name, email, age, activo FROM users WHERE id=$1", id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"Error:": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, u)
}

// Create user

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	err := config.DB.QueryRow(
		"INSERT INTO users (name, email, age, activo) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Name, u.Email, u.Age, u.Activo,
	).Scan(&u.ID)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}

	respondJSON(w, 201, u)
}

// Update

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	_, err := config.DB.Exec(
		"UPDATE users SET name=$1, email=$2, age=$3, activo=$4 WHERE id=$5",
		u.Name, u.Email, u.Age, u.Activo, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato Actualizado"})
}

// Delete

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM users WHERE id=$1", id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}