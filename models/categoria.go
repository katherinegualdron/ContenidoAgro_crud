package models

import "time"

type Categoria struct {
	IdCategoria       int       `json:"id_categoria"`
	NombreCategoria   string    `json:"nombre_categoria"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
