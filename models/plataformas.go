package models

import "time"

type Plataformas struct {
	ID            int       `db:"id" json:"id"`
	SucursalID    int       `db:"sucursal_id" json:"sucursal_id"`
	Nombre        string    `db:"nombre" json:"nombre"`
	FechaCreacion time.Time `db:"fecha_creacion" json:"fecha_creacion"`
}
