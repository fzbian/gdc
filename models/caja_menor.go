package models

type CajaMenor struct {
	SucursalID int `db:"sucursal_id" json:"sucursal_id"`
	Saldo      int `db:"saldo" json:"saldo"`
}
