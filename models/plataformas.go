package models

type Plataformas struct {
	ID            int    `db:"id" json:"id"`
	SucursalID    int    `db:"sucursal_id" json:"sucursal_id"`
	Nombre        string `db:"nombre" json:"nombre"`
	Saldo         int    `db:"saldo" json:"saldo"`
	FechaCreacion string `db:"fecha_creacion" json:"fecha_creacion"`
}
