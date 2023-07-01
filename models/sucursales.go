package models

import "time"

type Sucursal struct {
	ID            int       `db:"id" json:"id"`
	Nombre        string    `db:"nombre" json:"nombre"`
	FechaCreacion time.Time `db:"fecha_creacion" json:"fecha_creacion"`
}
