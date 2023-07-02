package movimientos_caja

import (
	"julio/controllers/cajeros"
	"julio/db"
	"julio/models"
)

func SaldoEsIgualAlSaldoDelCajero(saldoSalida, cajero_id int) bool {
	// Verificar que el saldo no sea menor a 0
	if saldoSalida < 0 {
		return false
	}
	// Verificar que el cajero exista
	if !cajeros.CajeroExistePorLaId(cajero_id) {
		return false
	}

	// Verificar que el saldo sea igual al saldo del cajero
	var c models.Cajero
	result := db.Db.Table("cajeros").Where("id = ?", cajero_id).First(&c)
	if result.Error != nil {
		return false
	}
	if c.Saldo != saldoSalida {
		return false
	}
	return true
}
