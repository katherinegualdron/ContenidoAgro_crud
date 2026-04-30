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

func ObtenerTemasForo(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_tema, titulo, descripcion, id_autor, id_categoria, estado, activo, fecha_creacion, fecha_modificacion FROM "TemaForo" ORDER BY id_tema`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar tema_foro")
		return
	}
	defer rows.Close()
	items := []models.TemaForo{}
	for rows.Next() {
		var item models.TemaForo
		if err := rows.Scan(&item.IdTema, &item.Titulo, &item.Descripcion, &item.IdAutor, &item.IdCategoria, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer tema_foro")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerTemaForoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.TemaForo
	err = config.DB.QueryRow(`SELECT id_tema, titulo, descripcion, id_autor, id_categoria, estado, activo, fecha_creacion, fecha_modificacion FROM "TemaForo" WHERE id_tema = $1`, id).
		Scan(&item.IdTema, &item.Titulo, &item.Descripcion, &item.IdAutor, &item.IdCategoria, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "tema_foro no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar tema_foro")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearTemaForo(w http.ResponseWriter, r *http.Request) {
	var item models.TemaForo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarTemaForo(item.IdAutor, item.IdCategoria); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`INSERT INTO "TemaForo" (titulo, descripcion, id_autor, id_categoria, estado, activo) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_tema, titulo, descripcion, id_autor, id_categoria, estado, activo, fecha_creacion, fecha_modificacion`,
		item.Titulo, item.Descripcion, item.IdAutor, item.IdCategoria, item.Estado, item.Activo)
	err := row.Scan(&item.IdTema, &item.Titulo, &item.Descripcion, &item.IdAutor, &item.IdCategoria, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			writeError(w, http.StatusBadRequest, "estado no valido")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear tema_foro")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarTemaForo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.TemaForo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarTemaForo(item.IdAutor, item.IdCategoria); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`UPDATE "TemaForo" SET titulo = $1, descripcion = $2, id_autor = $3, id_categoria = $4, estado = $5, activo = $6 WHERE id_tema = $7 RETURNING id_tema, titulo, descripcion, id_autor, id_categoria, estado, activo, fecha_creacion, fecha_modificacion`,
		item.Titulo, item.Descripcion, item.IdAutor, item.IdCategoria, item.Estado, item.Activo, id)
	err = row.Scan(&item.IdTema, &item.Titulo, &item.Descripcion, &item.IdAutor, &item.IdCategoria, &item.Estado, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "tema_foro no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			writeError(w, http.StatusBadRequest, "estado no valido")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar tema_foro")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarTemaForo(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoContenido(w, r, `"TemaForo"`, "id_tema", "tema_foro")
}

func validarTemaForo(idAutor int, idCategoria int) error {
	if ok, err := existeRegistro(`"Usuarios"."Usuario"`, "id_usuario", idAutor); err != nil {
		return errors.New("error al validar id_autor")
	} else if !ok {
		return errors.New("id_autor no existe")
	}
	if ok, err := existeRegistro(`"Categoria"`, "id_categoria", idCategoria); err != nil {
		return errors.New("error al validar id_categoria")
	} else if !ok {
		return errors.New("id_categoria no existe")
	}
	return nil
}
