package transacciones

import (
	"errors"
	"fmt"
	"julio/controllers/cajeros"
	"julio/controllers/usuarios"
	"julio/db"
	"julio/models"
	"julio/utils"
	"time"
)

// CrearTransacciones
func CrearTransaccion(cajero_id, usuario_id, valor int, tipo, descripcion string, billetera bool) (string, error) {
	// Verificar que el cajero exista
	if !cajeros.CajeroExistePorLaId(cajero_id) {
		return "El cajero no existe", nil
	}
	// Verificar que la billetera sea false o true
	if billetera != false && billetera != true {
		return "La billetera no es valida", nil
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
		if billetera {
			// Sumarle 1000 al valor final
			valor += 1000
		}
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
		if billetera {
			// Si el valor esta entre el rango de 16000 y 300000
			if valor >= 160000 && valor <= 300000 {
				// Sumarle 1000 al valor final
				valor += 1000
			}
			// Si el valor esta entre el rango de 300001 y 499999
			if valor >= 300001 && valor <= 499999 {
				// Sumarle 2000 al valor final
				valor += 2000
			}
			// Si el valor esta entre el rango de 500000 y 799999
			if valor >= 500000 && valor <= 799999 {
				// Sumarle 3000 al valor final
				valor += 3000
			}
			// Si el valor esta entre el rango de 800000 y 1000000
			if valor >= 800000 && valor <= 1000000 {
				// Sumarle 4000 al valor final
				valor += 4000
			}
			// Si el valor es de mas de 1000000
			if valor > 1000000 {
				// Sumarle 5000 al valor final
				valor += 5000
			}
		}
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
func EditarTransaccion(usuario_id, transaccion_id int, tipo string, valor any) (string, error) {
	// Verificar que el usuario tenga el rango adecuado
	if !usuarios.RangoCorrecto(usuario_id) {
		return "", errors.New("El usuario no tiene el rango adecuado para realizar esta accion")
	}
	// Verificar que el id no sea vacio o menor a 0
	if transaccion_id <= 0 {
		return "", errors.New("El id no puede ser vacio o menor a 0")
	}
	// Verificar que la transaccion exista
	if !TransaccionExistePorLaId(transaccion_id) {
		return "", errors.New("La transaccion no existe")
	}
	// Verificar que el tipo no sea vacio
	Tipos := []string{"descripcion", "valor"}
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
	} else if tipo == "valor" {
		// Verificar que el valor no sea vacio o menor a 0
		if valor.(int) <= 0 {
			return "", errors.New("El valor no puede ser vacio o menor a 0")
		}
		// Actualizar el valor de la transaccion
		result := db.Db.Table("transacciones").Where("id = ?", transaccion_id).Update("valor", valor)
		if result.Error != nil {
			return "", result.Error
		}
	}
	// Modificar fecha_actualizacion
	result := db.Db.Table("transacciones").Where("id = ?", transaccion_id).Update("fecha_actualizacion", time.Now())
	if result.Error != nil {
		return "", result.Error
	}
	return "Transaccion editada", nil
}

// EnlistarTransacciones devuelve todas las transacciones en un [][]string
func EnlistarTransacciones() ([][]string, error) {
	// Obtener todas las transacciones de la tabla
	var transacciones []models.Transacciones
	if err := db.Db.Find(&transacciones).Error; err != nil {
		return nil, err
	}

	// Construir la matriz de resultados
	resultados := make([][]string, len(transacciones)+1)
	resultados[0] = []string{"ID", "UsuarioID", "CajeroID", "Tipo", "Descripcion", "Billetera", "Valor", "FechaCreacion", "FechaActualizacion"}

	for i, transaccion := range transacciones {
		resultados[i+1] = []string{
			fmt.Sprintf("%d", transaccion.CajeroID),
			fmt.Sprintf("%s", usuarios.NombrePorLaId(transaccion.UsuarioID)),
			transaccion.Tipo,
			transaccion.Descripcion,
			fmt.Sprintf("%t", transaccion.Billetera),
			fmt.Sprintf("%d", transaccion.Valor),
			//fmt.Sprintf("%s", transaccion.FechaCreacion.Format(time.RFC3339)),
			//fmt.Sprintf("%s", transaccion.FechaActualizacion.Format(time.RFC3339)),
		}
	}

	return resultados, nil
}
