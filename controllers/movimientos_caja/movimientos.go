package movimientos_caja

import (
	"errors"
	"fmt"
	"julio/controllers/cajeros"
	"julio/controllers/sucursales"
	"julio/controllers/usuarios"
	"julio/db"
	"julio/utils"
	"time"
)

// CrearSesion
func CrearSesion(usuario_id, cajero_id, sucursal_id, clave, entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50 int) (string, error) {
	// Verificar que el usuario exista
	if !usuarios.UsuarioExistePorLaId(usuario_id) {
		return "", errors.New("El usuario no existe")
	}
	// Verificar que el cajero exista
	if !cajeros.CajeroExistePorLaId(cajero_id) {
		return "", errors.New("El cajero no existe")
	}
	// Verificar que el usuario no este en un cajero distinto
	if usuarios.UsuarioEnSesion(usuario_id) {
		return "", errors.New("El usuario ya esta en sesion")
	}
	// Verificar que el cajero pertenezca a la sucursal
	if !cajeros.CajeroPerteneceASucursal(cajero_id, sucursal_id) {
		return "", errors.New("El cajero no pertenece a la sucursal")
	}
	// Verificar que la sucursal exista
	if !sucursales.SucursalExistePorLaId(sucursal_id) {
		return "", errors.New("La sucursal no existe")
	}
	// Verificar que la clave sea correcta
	if !usuarios.ClaveCorrecta(usuario_id, clave) {
		return "", errors.New("La clave es incorrecta")
	}

	// Asignar usuario a cajero
	result := db.Db.Table("cajeros").Where("id = ?", cajero_id).Update("usuario_id", usuario_id)
	if result.Error != nil {
		return "", result.Error
	}
	// Actualizar usuario
	result = db.Db.Table("usuarios").Where("id = ?", usuario_id).Update("usuario_en_sesion", true)
	if result.Error != nil {
		return "", result.Error
	}

	// TODO: Pasar a gorm
	result = db.Db.Exec("INSERT INTO movimientos_caja (cajero_id, usuario_id, fecha_entrada, entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", cajero_id, usuario_id, time.Now(), entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50)
	if result.Error != nil {
		return "", result.Error
	}

	saldos, err := utils.SumarSaldos(entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50)
	if err != nil {
		return "", err
	}

	// Actualizar saldo en cajero
	result = db.Db.Table("cajeros").Where("id = ?", cajero_id).Update("saldo", saldos)
	if result.Error != nil {
		return "", result.Error
	}
	statusMessage := fmt.Sprintf("Sesion creada con un saldo de %d pesos", saldos)
	return statusMessage, nil
}

// CerrarSesion
func CerrarSesion(usuario_id, cajero_id, sucursal_id, salida_billetes_100, salida_billetes_50, salida_billetes_20, salida_billetes_10, salida_billetes_5, salida_billetes_2, salida_billetes_1, salida_monedas_1000, salida_monedas_500, salida_monedas_200, salida_monedas_100, salida_monedas_50 int) (string, error) {
	// Verificar que el usuario exista
	if !usuarios.UsuarioExistePorLaId(usuario_id) {
		return "", errors.New("El usuario no existe")
	}
	// Verificar que el cajero exista
	if !cajeros.CajeroExistePorLaId(cajero_id) {
		return "", errors.New("El cajero no existe")
	}
	// Verificar que el usuario este en el cajero
	if !cajeros.UsuarioEstaEnElCajero(cajero_id, usuario_id) {
		return "", errors.New("El usuario no esta en el cajero")
	}
	// Verificar que la sucursal exista
	if !sucursales.SucursalExistePorLaId(sucursal_id) {
		return "", errors.New("La sucursal no existe")
	}

	saldosSalida, err := utils.SumarSaldos(salida_billetes_100, salida_billetes_50, salida_billetes_20, salida_billetes_10, salida_billetes_5, salida_billetes_2, salida_billetes_1, salida_monedas_1000, salida_monedas_500, salida_monedas_200, salida_monedas_100, salida_monedas_50)
	if err != nil {
		return "", err
	}
	// Verificar que el saldo de salida sea el mismo que esta en el cajero
	if SaldoEsIgualAlSaldoDelCajero(saldosSalida, cajero_id) {
		return "", errors.New("El saldo de salida no es igual al saldo del cajero")
	}

	// Actualizar usuario quitandole el ensesion
	result := db.Db.Table("usuarios").Where("id = ?", usuario_id).Update("usuario_en_sesion", false)
	if result.Error != nil {
		return "", result.Error
	}
	// Actualizar cajero quitandole el usuario
	result = db.Db.Table("cajeros").Where("id = ?", cajero_id).Update("usuario_id", nil)
	if result.Error != nil {
		return "", result.Error
	}
	// Actualizar el movimiento y agregarle los salidas
	result = db.Db.Exec("UPDATE movimientos_caja SET fecha_salida = ?, salida_billetes_100 = ?, salida_billetes_50 = ?, salida_billetes_20 = ?, salida_billetes_10 = ?, salida_billetes_5 = ?, salida_billetes_2 = ?, salida_billetes_1 = ?, salida_monedas_1000 = ?, salida_monedas_500 = ?, salida_monedas_200 = ?, salida_monedas_100 = ?, salida_monedas_50 = ? WHERE cajero_id = ? AND usuario_id = ? AND fecha_salida IS NULL", time.Now(), salida_billetes_100, salida_billetes_50, salida_billetes_20, salida_billetes_10, salida_billetes_5, salida_billetes_2, salida_billetes_1, salida_monedas_1000, salida_monedas_500, salida_monedas_200, salida_monedas_100, salida_monedas_50, cajero_id, usuario_id)
	if result.Error != nil {
		return "", result.Error
	}

	statusMessage := fmt.Sprintf("Sesion cerrada con un saldo de %d pesos", saldosSalida)
	return statusMessage, nil
}
