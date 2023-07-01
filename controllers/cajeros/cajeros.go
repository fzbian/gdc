package cajeros

import (
	"errors"
	"julio/controllers/sucursales"
	"julio/controllers/usuarios"
	"julio/db"
	"julio/models"
	"julio/utils"
	"time"
)

// Crear cajero
func CrearCajero(sucursal_id, saldo int) (string, error) {
	// Verificar que la sucursal_id no sea 0 o menor a 0
	if sucursal_id <= 0 {
		return "", errors.New("La sucursal_id no es válida")
	}
	// Verificar que la sucursal exista
	if !sucursales.SucursalExistePorLaId(sucursal_id) {
		return "", errors.New("La sucursal no existe")
	}
	// Verificar que el saldo no sea menor a 0
	if saldo < 0 {
		return "", errors.New("El saldo no puede ser menor a 0")
	}

	// Crear el cajero
	result := db.Db.Table("cajeros").Create(&models.Cajero{
		SucursalID:    sucursal_id,
		UsuarioID:     nil,
		Saldo:         saldo,
		FechaCreacion: time.Now(),
	})
	if result.Error != nil {
		return "", result.Error
	}
	return "Cajero creado", nil
}

// Eliminar cajero
func EliminarCajero(id int) (string, error) {
	// Verificar que la id no sea 0 o menor a 0
	if id <= 0 {
		return "", errors.New("La id no es válida")
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(id) {
		return "", errors.New("El cajero no existe")
	}

	// Eliminar el cajero
	result := db.Db.Table("cajeros").Where("id = ?", id).Delete(&models.Cajero{})
	if result.Error != nil {
		return "", result.Error
	}
	return "Cajero eliminado", nil
}

// Editar cajero
func EditarCajero(id int, tipo string, valor any) (string, error) {
	// Verificar que la id no sea 0 o menor a 0
	if id <= 0 {
		return "", errors.New("La id no es válida")
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(id) {
		return "", errors.New("El cajero no existe")
	}
	// Verificar que el tipo no este vacio
	if tipo == "" {
		return "", errors.New("El tipo no puede estar vacio")
	}
	// Verificar que el valor no este vacio
	if valor == "" {
		return "", errors.New("El valor no puede estar vacio")
	}
	// Verificar que el valor sea valido
	TiposCajeros := []string{"sucursal_id", "usuario_id", "saldo"}
	if !utils.Contains(TiposCajeros, tipo) {
		return "", errors.New("El tipo no es valido")
	}
	if tipo == "sucursal_id" {
		// Verificar que la sucursal_id no sea 0 o menor a 0
		if valor.(int) <= 0 {
			return "", errors.New("La sucursal_id no es válida")
		}
		// Verificar que la sucursal exista
		if !sucursales.SucursalExistePorLaId(valor.(int)) {
			return "", errors.New("La sucursal no existe")
		}
	} else if tipo == "usuario_id" {
		// Verificar que el usuario_id no sea 0 o menor a 0
		if valor.(int) <= 0 {
			return "", errors.New("El usuario_id no es válido")
		}
		// Verificar que el usuario exista
		if !usuarios.UsuarioExistePorLaId(valor.(int)) {
			return "", errors.New("El usuario no existe")
		}
		// Verificar que el usuario no tenga un cajero ya asignado
		if UsuarioTieneCajero(valor.(int)) {
			return "", errors.New("El usuario ya tiene un cajero asignado")
		}
	} else if tipo == "saldo" {
		// Verificar que el saldo no sea menor a 0
		if valor.(int) < 0 {
			return "", errors.New("El saldo no puede ser menor a 0")
		}
	}

	// Editar el cajero
	result := db.Db.Table("cajeros").Where("id = ?", id).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Cajero editado", nil
}

// Enlistar cajeros
func EnlistarCajeros() ([]models.Cajero, error) {
	// Enlistar los cajeros
	var cajeros []models.Cajero
	result := db.Db.Table("cajeros").Find(&cajeros)
	if result.Error != nil {
		return nil, result.Error
	}
	return cajeros, nil
}
