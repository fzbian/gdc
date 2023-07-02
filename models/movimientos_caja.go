package models

import "time"

type MovimientosCaja struct {
	ID                 int        `db:"id" json:"id" gorm:"column:id"`
	CajeroID           int        `db:"cajero_id" json:"cajero_id" gorm:"cajero_id" gorm:"column:cajero_id"`
	UsuarioID          int        `db:"usuario_id" json:"usuario_id" gorm:"usuario_id" gorm:"column:usuario_id"`
	FechaEntrada       time.Time  `db:"fecha_entrada" json:"fecha_entrada" gorm:"fecha_entrada" gorm:"column:fecha_entrada"`
	EntradaBilletes100 *int       `db:"entrada_billetes_100" json:"entrada_billetes_100" gorm:"entrada_billetes_100" gorm:"column:entrada_billetes_100"`
	EntradaBilletes50  *int       `db:"entrada_billetes_50" json:"entrada_billetes_50" gorm:"entrada_billetes_50" gorm:"column:entrada_billetes_50"`
	EntradaBilletes20  *int       `db:"entrada_billetes_20" json:"entrada_billetes_20" gorm:"entrada_billetes_20" gorm:"column:entrada_billetes_20"`
	EntradaBilletes10  *int       `db:"entrada_billetes_10" json:"entrada_billetes_10" gorm:"entrada_billetes_10" gorm:"column:entrada_billetes_10"`
	EntradaBilletes5   *int       `db:"entrada_billetes_5" json:"entrada_billetes_5" gorm:"entrada_billetes_5" gorm:"column:entrada_billetes_5"`
	EntradaBilletes2   *int       `db:"entrada_billetes_2" json:"entrada_billetes_2" gorm:"entrada_billetes_2" gorm:"column:entrada_billetes_2"`
	EntradaBilletes1   *int       `db:"entrada_billetes_1" json:"entrada_billetes_1" gorm:"entrada_billetes_1" gorm:"column:entrada_billetes_1"`
	EntradaMonedas1000 *int       `db:"entrada_monedas_1000" json:"entrada_monedas_1000" gorm:"entrada_monedas_1000" gorm:"column:entrada_monedas_1000"`
	EntradaMonedas500  *int       `db:"entrada_monedas_500" json:"entrada_monedas_500" gorm:"entrada_monedas_500" gorm:"column:entrada_monedas_500"`
	EntradaMonedas200  *int       `db:"entrada_monedas_200" json:"entrada_monedas_200" gorm:"entrada_monedas_200" gorm:"column:entrada_monedas_200"`
	EntradaMonedas100  *int       `db:"entrada_monedas_100" json:"entrada_monedas_100" gorm:"entrada_monedas_100" gorm:"column:entrada_monedas_100"`
	EntradaMonedas50   *int       `db:"entrada_monedas_50" json:"entrada_monedas_50" gorm:"entrada_monedas_50" gorm:"column:entrada_monedas_50"`
	SalidaBilletes100  *int       `db:"salida_billetes_100" json:"salida_billetes_100" gorm:"salida_billetes_100" gorm:"column:salida_billetes_100"`
	SalidaBilletes50   *int       `db:"salida_billetes_50" json:"salida_billetes_50" gorm:"salida_billetes_50" gorm:"column:salida_billetes_50"`
	SalidaBilletes20   *int       `db:"salida_billetes_20" json:"salida_billetes_20" gorm:"salida_billetes_20" gorm:"column:salida_billetes_20"`
	SalidaBilletes10   *int       `db:"salida_billetes_10" json:"salida_billetes_10" gorm:"salida_billetes_10" gorm:"column:salida_billetes_10"`
	SalidaBilletes5    *int       `db:"salida_billetes_5" json:"salida_billetes_5" gorm:"salida_billetes_5" gorm:"column:salida_billetes_5"`
	SalidaBilletes2    *int       `db:"salida_billetes_2" json:"salida_billetes_2" gorm:"salida_billetes_2" gorm:"column:salida_billetes_2"`
	SalidaBilletes1    *int       `db:"salida_billetes_1" json:"salida_billetes_1" gorm:"salida_billetes_1" gorm:"column:salida_billetes_1"`
	SalidaMonedas1000  *int       `db:"salida_monedas_1000" json:"salida_monedas_1000" gorm:"salida_monedas_1000" gorm:"column:salida_monedas_1000"`
	SalidaMonedas500   *int       `db:"salida_monedas_500" json:"salida_monedas_500" gorm:"salida_monedas_500" gorm:"column:salida_monedas_500"`
	SalidaMonedas200   *int       `db:"salida_monedas_200" json:"salida_monedas_200" gorm:"salida_monedas_200" gorm:"column:salida_monedas_200"`
	SalidaMonedas100   *int       `db:"salida_monedas_100" json:"salida_monedas_100" gorm:"salida_monedas_100"  gorm:"column:salida_monedas_100"`
	SalidaMonedas50    *int       `db:"salida_monedas_50" json:"salida_monedas_50" gorm:"salida_monedas_50" gorm:"column:salida_monedas_50"`
	FechaSalida        *time.Time `db:"fecha_salida" json:"fecha_salida" gorm:"fecha_salida" gorm:"column:fecha_salida"`
}
