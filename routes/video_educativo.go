package models

import "time"

type VideoEducativo struct {
	IdVideo           int       `json:"id_video"`
	Titulo            string    `json:"titulo"`
	Descripcion       string    `json:"descripcion"`
	UrlVideo          string    `json:"url_video"`
	IdUsuario         int       `json:"id_usuario"`
	Estado            string    `json:"estado"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
