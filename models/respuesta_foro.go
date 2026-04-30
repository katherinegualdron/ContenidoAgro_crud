package models

import "time"

type RespuestaForo struct {
	IdRespuesta       int       `json:"id_respuesta"`
	IdTema            int       `json:"id_tema"`
	IdAutor           int       `json:"id_autor"`
	Descripcion       string    `json:"descripcion"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
