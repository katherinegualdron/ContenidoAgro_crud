package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"CONTENIDO/config"
	"CONTENIDO/models"
)

func ObtenerRespuestasForo(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_respuesta, id_tema, id_autor, descripcion, activo, fecha_creacion, fecha_modificacion FROM "RespuestaForo" ORDER BY id_respuesta`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar respuesta_foro")
		return
	}
	defer rows.Close()
	items := []models.RespuestaForo{}
	for rows.Next() {
		var item models.RespuestaForo
		if err := rows.Scan(&item.IdRespuesta, &item.IdTema, &item.IdAutor, &item.Descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer respuesta_foro")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerRespuestaForoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.RespuestaForo
	err = config.DB.QueryRow(`SELECT id_respuesta, id_tema, id_autor, descripcion, activo, fecha_creacion, fecha_modificacion FROM "RespuestaForo" WHERE id_respuesta = $1`, id).
		Scan(&item.IdRespuesta, &item.IdTema, &item.IdAutor, &item.Descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "respuesta_foro no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar respuesta_foro")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearRespuestaForo(w http.ResponseWriter, r *http.Request) {
	var item models.RespuestaForo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRespuestaForo(item.IdTema, item.IdAutor); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := config.DB.QueryRow(`INSERT INTO "RespuestaForo" (id_tema, id_autor, descripcion, activo) VALUES ($1, $2, $3, $4) RETURNING id_respuesta, id_tema, id_autor, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.IdTema, item.IdAutor, item.Descripcion, item.Activo).
		Scan(&item.IdRespuesta, &item.IdTema, &item.IdAutor, &item.Descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear respuesta_foro")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarRespuestaForo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.RespuestaForo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRespuestaForo(item.IdTema, item.IdAutor); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = config.DB.QueryRow(`UPDATE "RespuestaForo" SET id_tema = $1, id_autor = $2, descripcion = $3, activo = $4 WHERE id_respuesta = $5 RETURNING id_respuesta, id_tema, id_autor, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.IdTema, item.IdAutor, item.Descripcion, item.Activo, id).
		Scan(&item.IdRespuesta, &item.IdTema, &item.IdAutor, &item.Descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "respuesta_foro no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar respuesta_foro")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarRespuestaForo(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoContenido(w, r, `"RespuestaForo"`, "id_respuesta", "respuesta_foro")
}

func validarRespuestaForo(idTema int, idAutor int) error {
	if ok, err := existeRegistro(`"TemaForo"`, "id_tema", idTema); err != nil {
		return errors.New("error al validar id_tema")
	} else if !ok {
		return errors.New("id_tema no existe")
	}
	if ok, err := existeRegistro(`"Usuarios"."Usuario"`, "id_usuario", idAutor); err != nil {
		return errors.New("error al validar id_autor")
	} else if !ok {
		return errors.New("id_autor no existe")
	}
	return nil
}
