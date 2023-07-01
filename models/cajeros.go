package models

import "time"

type Cajero struct {
	ID            int       `db:"id" json:"id"`
	SucursalID    int       `db:"sucursal_id" json:"sucursal_id"`
	UsuarioID     *int      `db:"usuario_id" json:"usuario_id"`
	Saldo         int       `db:"saldo" json:"saldo"`
	FechaCreacion time.Time `db:"fecha_creacion" json:"fecha_creacion"`
}
