package cajeros

import (
	"errors"
	"julio/db"
	"julio/models"
)

// CajeroExistePorLaId
func CajeroExistePorLaId(id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var s models.Cajero
	result := db.Db.Table("Cajeros").Where("id = ?", id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

// Verificar que un usuario no tenga un cajero asignado,
func UsuarioTieneCajero(usuario_id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if usuario_id <= 0 {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var s models.Cajero
	result := db.Db.Table("cajeros").Where("usuario_id = ?", usuario_id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

// Devolver el saldo del cajero por la id
func DevolverSaldoCajeroPorLaId(id int) (int, error) {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return 0, errors.New("La id no es válida")
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(id) {
		return 0, errors.New("El cajero no existe")
	}

	var s models.Cajero
	result := db.Db.Table("cajeros").Where("id = ?", id).First(&s)
	if result.Error != nil {
		return 0, result.Error
	}
	return s.Saldo, nil
}
