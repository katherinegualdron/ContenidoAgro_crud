package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"CONTENIDO/config"
	"CONTENIDO/models"

	"github.com/lib/pq"
)

func ObtenerCategorias(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_categoria, nombre_categoria, activo, fecha_creacion, fecha_modificacion FROM "Categoria" ORDER BY id_categoria`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar categoria")
		return
	}
	defer rows.Close()
	items := []models.Categoria{}
	for rows.Next() {
		var item models.Categoria
		if err := rows.Scan(&item.IdCategoria, &item.NombreCategoria, &item.Activo, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer categoria")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerCategoriaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Categoria
	err = config.DB.QueryRow(`SELECT id_categoria, nombre_categoria, activo, fecha_creacion, fecha_modificacion FROM "Categoria" WHERE id_categoria = $1`, id).
		Scan(&item.IdCategoria, &item.NombreCategoria, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "categoria no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar categoria")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var item models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	err := config.DB.QueryRow(`INSERT INTO "Categoria" (nombre_categoria, activo) VALUES ($1, $2) RETURNING id_categoria, nombre_categoria, activo, fecha_creacion, fecha_modificacion`,
		item.NombreCategoria, item.Activo).Scan(&item.IdCategoria, &item.NombreCategoria, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "nombre_categoria ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear categoria")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	err = config.DB.QueryRow(`UPDATE "Categoria" SET nombre_categoria = $1, activo = $2 WHERE id_categoria = $3 RETURNING id_categoria, nombre_categoria, activo, fecha_creacion, fecha_modificacion`,
		item.NombreCategoria, item.Activo, id).Scan(&item.IdCategoria, &item.NombreCategoria, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "categoria no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "nombre_categoria ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar categoria")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarCategoria(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoContenido(w, r, `"Categoria"`, "id_categoria", "categoria")
}
