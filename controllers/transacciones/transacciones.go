package transacciones

import (
	"errors"
	"julio/controllers/cajeros"
	"julio/controllers/usuarios"
	"julio/db"
	"julio/models"
	"julio/utils"
	"time"
)

// CrearTransacciones
func CrearTransaccion(cajero_id, usuario_id, valor int, tipo, descripcion string) (string, error) {
	// Verificar que el cajero exista
	if !cajeros.CajeroExistePorLaId(cajero_id) {
		return "El cajero no existe", nil
	}
	// Verificar que el usuario exista
	if !usuarios.UsuarioExistePorLaId(usuario_id) {
		return "El usuario no existe", nil
	}
	// Verificar que el usuario este en el cajero
	if !cajeros.UsuarioEstaEnElCajero(usuario_id, cajero_id) {
		return "El usuario no esta en el cajero", nil
	}
	// Verificar que el valor no sea menor a 0
	if valor <= 0 {
		return "El valor no puede ser menor a 0", nil
	}
	// Verificar que el tipo no sea vacio
	Tipos := []string{"DEPOSITO", "RETIRO"}
	if !utils.Contains(Tipos, tipo) {
		return "El tipo no es valido", nil
	}
	// Si la descripcion es vacia, asignarle un valor por defecto
	if descripcion == "" {
		descripcion = "Sin descripcion"
	}
	// Si la transaccion es un DEPOSITO, sumar el valor al saldo del cajero
	if tipo == "DEPOSITO" {
		// Obtener el saldo del cajero
		saldo, err := cajeros.DevolverSaldoCajeroPorLaId(cajero_id)
		if err != nil {
			return "", err
		}
		// Sumar el valor al saldo del cajero
		saldo += valor
		// Actualizar el saldo del cajero
		result := db.Db.Table("cajeros").Where("id = ?", cajero_id).Update("saldo", saldo)
		if result.Error != nil {
			return "", result.Error
		}
	} else if tipo == "RETIRO" {
		// Si la transaccion es un RETIRO, restar el valor al saldo del cajero
		// Obtener el saldo del cajero
		saldo, err := cajeros.DevolverSaldoCajeroPorLaId(cajero_id)
		if err != nil {
			return "", err
		}
		// Si el cajero no tiene el saldo suficiente para entregar envia un error
		if saldo < valor {
			return "", errors.New("El cajero no tiene el saldo suficiente para entregar")
		}
		// Restar el valor al saldo del cajero
		saldo -= valor
		// Actualizar el saldo del cajero
		result := db.Db.Table("cajeros").Where("id = ?", cajero_id).Update("saldo", saldo)
		if result.Error != nil {
			return "", result.Error
		}
	}
	// Crear la transaccion
	result := db.Db.Table("transacciones").Create(&models.Transacciones{
		CajeroID:      cajero_id,
		UsuarioID:     usuario_id,
		Tipo:          tipo,
		Descripcion:   descripcion,
		Valor:         valor,
		FechaCreacion: time.Now(),
	})
	if result.Error != nil {
		return "", result.Error
	}
	return "Transaccion creada", nil
}

// EliminarTransaccion
func EliminarTransaccion(id int) (string, error) {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return "", errors.New("El id no puede ser vacio o menor a 0")
	}
	// Verificar que la transaccion exista
	if !TransaccionExistePorLaId(id) {
		return "", errors.New("La transaccion no existe")
	}

	// Eliminar la transaccion
	result := db.Db.Table("transacciones").Where("id = ?", id).Delete(&models.Transacciones{})
	if result.Error != nil {
		return "", result.Error
	}
	return "Transaccion eliminada", nil
}

// EditarTransaccion
func EditarTransaccion(transaccion_id int, tipo string, valor any) (string, error) {
	// Verificar que el id no sea vacio o menor a 0
	if transaccion_id <= 0 {
		return "", errors.New("El id no puede ser vacio o menor a 0")
	}
	// Verificar que la transaccion exista
	if !TransaccionExistePorLaId(transaccion_id) {
		return "", errors.New("La transaccion no existe")
	}
	// Verificar que el tipo no sea vacio
	Tipos := []string{"descripcion"}
	if !utils.Contains(Tipos, tipo) {
		return "El tipo no es valido", nil
	}
	if tipo == "descripcion" {
		// Verificar que la descripcion no sea vacia
		if valor.(string) == "" {
			descripcion := "Sin descripcion"
			valor = descripcion
		}
		// Actualizar la descripcion de la transaccion
		result := db.Db.Table("transacciones").Where("id = ?", transaccion_id).Update("descripcion", valor)
		if result.Error != nil {
			return "", result.Error
		}
	}
	return "Transaccion editada", nil
}

// EnlistarTransacciones
func EnlistarTransacciones() ([]models.Transacciones, error) {
	// Enlistar transacciones
	var transacciones []models.Transacciones
	result := db.Db.Table("cajeros").Find(&transacciones)
	if result.Error != nil {
		return nil, result.Error
	}
	return transacciones, nil
}
