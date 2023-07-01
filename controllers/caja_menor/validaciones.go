package caja_menor

import (
	"julio/controllers/sucursales"
	"julio/db"
	"julio/models"
)

func CajaMenorExistePorLaSucursalId(id_sucursal int) bool {
	// Verificar que el id_sucursal no sea vacio o menor a 0
	if id_sucursal <= 0 {
		return false
	}
	// Verificar que la sucursal exista
	if !sucursales.SucursalExistePorLaId(id_sucursal) {
		return false
	}

	// Verificar que la caja menor exista
	var c models.CajaMenor
	result := db.Db.Table("caja_menor").Where("sucursal_id = ?", id_sucursal).First(&c)
	if result.Error != nil {
		return false
	}
	return true
}
