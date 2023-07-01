package sucursales

import (
	"julio/db"
	"julio/models"
)

// Verificar que la sucursal no exista
func SucursalExistePorElNombre(nombre string) bool {
	// Verificar que el nombre no sea vacio
	if nombre == "" {
		return false
	}
	// Verificar que el nombre no exista en la base de datos
	var s models.Sucursal
	result := db.Db.Table("sucursales").Where("nombre = ?", nombre).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

// SucursalExistePorLaId
func SucursalExistePorLaId(id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var s models.Sucursal
	result := db.Db.Table("sucursales").Where("id = ?", id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}
