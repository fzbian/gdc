package models

import "time"

type Transacciones struct {
	ID                 int       `db:"id" json:"id"`
	UsuarioID          int       `db:"usuario_id" json:"usuario_id"`
	CajeroID           int       `db:"cajero_id" json:"cajero_id"`
	Tipo               string    `db:"tipo" json:"tipo"`
	Descripcion        string    `db:"descripcion" json:"descripcion"`
	Billetera          bool      `db:"billetera" json:"billetera"`
	Valor              int       `db:"valor" json:"valor"`
	FechaCreacion      time.Time `db:"fecha_creacion" json:"fecha_creacion"`
	FechaActualizacion time.Time `db:"fecha_actualizacion" json:"fecha_actualizacion"`
}
