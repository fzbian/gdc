package models

import "time"

type Usuarios struct {
	Id              int       `db:"id" json:"id"`
	Usuario         string    `db:"usuario" json:"usuario"`
	Nombre          string    `db:"nombre" json:"nombre"`
	Rango           int       `db:"rango" json:"rango"`
	Clave           int       `db:"clave" json:"clave"`
	UsuarioEnSesion bool      `db:"usuario_en_sesion" json:"usuario_en_sesion"`
	FechaCreacion   time.Time `db:"fecha_creacion" json:"fecha_creacion"`
}
