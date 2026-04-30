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

func ObtenerVideosEducativos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_video, titulo, descripcion, url_video, id_usuario, estado, activo, fecha_creacion, fecha_modificacion FROM "VideoEducativo" ORDER BY id_video`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar video_educativo")
		return
	}
	defer rows.Close()
	items := []models.VideoEducativo{}
	for rows.Next() {
		var item models.VideoEducativo
		if err := rows.Scan(&item.IdVideo, &item.Titulo, &item.Descripcion, &item.UrlVideo, &item.IdUsuario, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer video_educativo")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerVideoEducativoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.VideoEducativo
	err = config.DB.QueryRow(`SELECT id_video, titulo, descripcion, url_video, id_usuario, estado, activo, fecha_creacion, fecha_modificacion FROM "VideoEducativo" WHERE id_video = $1`, id).
		Scan(&item.IdVideo, &item.Titulo, &item.Descripcion, &item.UrlVideo, &item.IdUsuario, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "video_educativo no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar video_educativo")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearVideoEducativo(w http.ResponseWriter, r *http.Request) {
	var item models.VideoEducativo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if ok, err := existeRegistro(`"Usuarios"."Usuario"`, "id_usuario", item.IdUsuario); err != nil {
		writeError(w, http.StatusBadRequest, "error al validar id_usuario")
		return
	} else if !ok {
		writeError(w, http.StatusBadRequest, "id_usuario no existe")
		return
	}
	row := config.DB.QueryRow(`INSERT INTO "VideoEducativo" (titulo, descripcion, url_video, id_usuario, estado, activo) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_video, titulo, descripcion, url_video, id_usuario, estado, activo, fecha_creacion, fecha_modificacion`,
		item.Titulo, item.Descripcion, item.UrlVideo, item.IdUsuario, item.Estado, item.Activo)
	err := row.Scan(&item.IdVideo, &item.Titulo, &item.Descripcion, &item.UrlVideo, &item.IdUsuario, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			writeError(w, http.StatusBadRequest, "estado no valido")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear video_educativo")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarVideoEducativo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.VideoEducativo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if ok, err := existeRegistro(`"Usuarios"."Usuario"`, "id_usuario", item.IdUsuario); err != nil {
		writeError(w, http.StatusBadRequest, "error al validar id_usuario")
		return
	} else if !ok {
		writeError(w, http.StatusBadRequest, "id_usuario no existe")
		return
	}
	row := config.DB.QueryRow(`UPDATE "VideoEducativo" SET titulo = $1, descripcion = $2, url_video = $3, id_usuario = $4, estado = $5, activo = $6 WHERE id_video = $7 RETURNING id_video, titulo, descripcion, url_video, id_usuario, estado, activo, fecha_creacion, fecha_modificacion`,
		item.Titulo, item.Descripcion, item.UrlVideo, item.IdUsuario, item.Estado, item.Activo, id)
	err = row.Scan(&item.IdVideo, &item.Titulo, &item.Descripcion, &item.UrlVideo, &item.IdUsuario, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "video_educativo no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			writeError(w, http.StatusBadRequest, "estado no valido")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar video_educativo")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarVideoEducativo(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoContenido(w, r, `"VideoEducativo"`, "id_video", "video_educativo")
}
