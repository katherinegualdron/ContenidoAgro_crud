package models

import "time"

type TemaForo struct {
	IdTema            int       `json:"id_tema"`
	Titulo            string    `json:"titulo"`
	Descripcion       string    `json:"descripcion"`
	IdAutor           int       `json:"id_autor"`
	IdCategoria       int       `json:"id_categoria"`
	Estado            string    `json:"estado"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
