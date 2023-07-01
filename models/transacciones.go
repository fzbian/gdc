package models

import "time"

type Transacciones struct {
	ID            int       `db:"id" json:"id"`
	CajeroID      int       `db:"cajero_id" json:"cajero_id"`
	Tipo          string    `db:"tipo" json:"tipo"`
	Descripcion   string    `db:"descripcion" json:"descripcion"`
	Valor         int       `db:"valor" json:"valor"`
	FechaCreacion time.Time `db:"fecha_creacion" json:"fecha_creacion"`
}
