package movimientos_caja

import (
	"julio/controllers/cajeros"
	"julio/db"
	"julio/models"
)

func SaldoEsIgualAlSaldoDelCajero(saldoSalida, cajero_id int) (*int, bool) {
	// Verificar que el saldo no sea menor a 0
	if saldoSalida < 0 {
		return nil, false
	}
	// Verificar que el cajero exista
	if !cajeros.CajeroExistePorLaId(cajero_id) {
		return nil, false
	}

	// Verificar que el saldo sea igual al saldo del cajero
	var c models.Cajero
	result := db.Db.Table("cajeros").Where("id = ?", cajero_id).First(&c)
	if result.Error != nil {
		return nil, false
	}
	// Si el saldo no es igual al saldo de salida retornar false y devolver de cuanto es el descruadre
	if c.Saldo != saldoSalida {
		descuadre := c.Saldo - saldoSalida
		return &descuadre, false
	}
	return nil, true
}
