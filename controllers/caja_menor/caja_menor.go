package caja_menor

import (
	"errors"
	"julio/controllers/cajeros"
	"julio/controllers/sucursales"
	"julio/db"
)

// Editar saldo de caja menor
func EditarSaldoCajaMenor(id_sucursal, id_cajero, saldo int) (string, error) {
	// Verificar que la id_sucursal no sea 0 o menor a 0
	if id_sucursal <= 0 {
		return "", errors.New("La id_sucursal no es válida")
	}
	// Verificar que el cajero exista
	if sucursales.SucursalExistePorLaId(id_sucursal) == false {
		return "", errors.New("Esta id_sucursal no pertenece a alguna sucursal")
	}
	// Verificar que el saldo no sea menor a 0
	if saldo < 0 {
		return "", errors.New("El saldo no puede ser menor a 0")
	}
	// Verificar que el cajero tenga saldo suficiente
	saldoCajero, err := cajeros.DevolverSaldoCajeroPorLaId(id_cajero)
	if err != nil {
		return "", err
	} else if saldoCajero < saldo {
		return "", errors.New("El saldo del cajero no es suficiente")
	}

	// Editar el saldo de la caja menor
	result := db.Db.Table("caja_menor").Where("sucursal_id = ?", id_sucursal).Update("saldo", saldo)
	if result.Error != nil {
		return "", result.Error
	}
	// Restar el saldo del cajero usando la funcion EditarCajero
	_, err = cajeros.EditarCajero(id_cajero, "saldo", saldoCajero-saldo)
	return "Saldo de caja menor editado", nil
}
