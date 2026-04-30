package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CONTENIDO/config"

	"github.com/gorilla/mux"
)

// Helper respuesta JSON
func responJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// get ALL
func GetALLCredentials(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, id_user, pssword FROM credentials")
	if err != nil {
		responJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Credential

	for rows.Next() {
		var c models.Credential
		if err := rows.Scan(&c.ID, &c.ID_User, &c.Password); err != nil {
			responJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		list = append(list, c)
	}

	if err := rows.Err(); err != nil {
		responJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	responJSON(w, 200, list)
}

// GetByID
func GetCredentialByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.Credential

	err := config.DB.QueryRow(
		"SELECT id, id_user, pssword FROM credentials WHERE id=$1", id,
	).Scan(&c.ID, &c.ID_User, &c.Password)
	if err != nil {
		responJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	responJSON(w, 200, c)
}

// create credential
func CreateCredential(w http.ResponseWriter, r *http.Request) {
	var c models.Credential
	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		"INSERT INTO credentials (id_user, pssword) VALUES ($1, $2) RETURNING id",
		c.ID_User, c.Password,
	).Scan(&c.ID)
	if err != nil {
		responJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	responJSON(w, 201, c)
}
func UpdateCredential(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.Credential
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		"UPDATE credentials SET pssword=$1 WHERE id=$2",
		c.Password, id,
	)

	if err != nil {
		responJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	responJSON(w, 200, map[string]string{"message": "Dato actualizado"})
}

// delete
func DeleteCredential(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM credentials WHERE id=$1", id)
	if err != nil {
		responJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	responJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}