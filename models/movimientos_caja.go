package models

import "time"

type MovimientoCaja struct {
	ID                 int
	CajeroID           int
	UsuarioID          int
	FechaEntrada       time.Time
	EntradaBilletes100 int
	EntradaBilletes50  int
	EntradaBilletes20  int
	EntradaBilletes10  int
	EntradaBilletes5   int
	EntradaBilletes2   int
	EntradaBilletes1   int
	EntradaMonedas1000 int
	EntradaMonedas500  int
	EntradaMonedas200  int
	EntradaMonedas100  int
	EntradaMonedas50   int
	SalidaBilletes100  int
	SalidaBilletes50   int
	SalidaBilletes20   int
	SalidaBilletes10   int
	SalidaBilletes5    int
	SalidaBilletes2    int
	SalidaBilletes1    int
	SalidaMonedas1000  int
	SalidaMonedas500   int
	SalidaMonedas200   int
	SalidaMonedas100   int
	SalidaMonedas50    int
	FechaSalida        time.Time
}
