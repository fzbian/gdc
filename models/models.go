package models

import "time"

type Sucursal struct {
	ID        int
	Nombre    string
	Direccion string
}

type Cajero struct {
	ID         int
	Nombre     string
	SucursalID int
}

type MovimientoCaja struct {
	ID             int
	CajeroID       int
	Fecha          time.Time
	DineroRecibido int64
	DineroSalida   int64
}

type Transaccion struct {
	ID         int
	CajeroID   int
	Tipo       string
	Plataforma string
	Referencia string
	Valor      int64
	Hora       time.Time
}
