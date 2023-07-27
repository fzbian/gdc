package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

/* Caja menor */

// Editar saldo de caja menor
func EditarSaldoCajaMenor(id_sucursal, id_cajero, saldo int) (string, error) {
	// Verificar que la id_sucursal no sea 0 o menor a 0
	if id_sucursal <= 0 {
		return "", errors.New("La id_sucursal no es válida")
	}
	// Verificar que el cajero exista
	if !SucursalExistePorLaId(id_sucursal) {
		return "", errors.New("Esta id_sucursal no pertenece a alguna sucursal")
	}
	// Verificar que el saldo no sea menor a 0
	if saldo < 0 {
		return "", errors.New("El saldo no puede ser menor a 0")
	}
	// Verificar que el cajero tenga saldo suficiente
	saldoCajero, err := DevolverSaldoCajeroPorLaId(id_cajero)
	if err != nil {
		return "", err
	} else if saldoCajero < saldo {
		return "", errors.New("El saldo del cajero no es suficiente")
	}

	// Editar el saldo de la caja menor
	result := Db.Table("caja_menor").Where("sucursal_id = ?", id_sucursal).Update("saldo", saldo)
	if result.Error != nil {
		return "", result.Error
	}
	// Restar el saldo del cajero usando la funcion EditarCajero
	_, err = EditarCajero(id_cajero, "saldo", saldoCajero-saldo)
	return "Saldo de caja menor editado", nil
}

// Devolver si la caja menor existe por la id de la sucursal
func CajaMenorExistePorLaSucursalId(id_sucursal int) bool {
	// Verificar que el id_sucursal no sea vacio o menor a 0
	if id_sucursal <= 0 {
		return false
	}
	// Verificar que la sucursal exista
	if !SucursalExistePorLaId(id_sucursal) {
		return false
	}

	// Verificar que la caja menor exista
	var c CajaMenor
	result := Db.Table("caja_menor").Where("sucursal_id = ?", id_sucursal).First(&c)
	if result.Error != nil {
		return false
	}
	return true
}

/* Cajeros */

// Crear cajero
func CrearCajero(sucursal_id int) (string, error) {
	// Verificar que la sucursal_id no sea 0 o menor a 0
	if sucursal_id <= 0 {
		return "", errors.New("La sucursal_id no es válida")
	}
	// Verificar que la sucursal exista
	if !SucursalExistePorLaId(sucursal_id) {
		return "", errors.New("La sucursal no existe")
	}

	// Crear el cajero
	result := Db.Table("cajeros").Create(&Cajero{
		SucursalID:    sucursal_id,
		UsuarioID:     nil,
		Saldo:         0,
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
	result := Db.Table("cajeros").Where("id = ?", id).Delete(&Cajero{})
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
	// Verificar que el valor no este vacio
	if valor == "" || valor == nil {
		return "", errors.New("El valor no puede estar vacio")
	}
	// Verificar que el valor sea valido
	TiposCajeros := []string{"sucursal_id", "usuario_id", "saldo"}
	if !Contains(TiposCajeros, tipo) {
		return "", errors.New("El tipo no es valido")
	}
	if tipo == "sucursal_id" {
		// Verificar que la sucursal_id no sea 0 o menor a 0
		if valor.(int) <= 0 {
			return "", errors.New("La sucursal_id no es válida")
		}
		// Verificar que la sucursal exista
		if !SucursalExistePorLaId(valor.(int)) {
			return "", errors.New("La sucursal no existe")
		}
	} else if tipo == "usuario_id" {
		// Verificar que el usuario_id no sea 0 o menor a 0
		if valor.(int) <= 0 {
			return "", errors.New("El usuario_id no es válido")
		}
		// Verificar que el usuario exista
		if !UsuarioExistePorLaId(valor.(int)) {
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
	println(fmt.Sprintf("ID: %d\nTipo: %s\nValor: %d", id, tipo, valor))
	result := Db.Table("cajeros").Where("id = ?", id).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Cajero editado", nil
}

// Enlistar cjaeros
func EnlistarCajeros() ([][]string, error) {
	// Obtener todas las cajeros de la tabla
	var cajeros []Cajero
	if err := Db.Find(&cajeros).Error; err != nil {
		return nil, err
	}

	// Construir la matriz de resultados
	resultados := make([][]string, len(cajeros)+1)
	//resultados[0] = []string{"ID", "UsuarioID", "SucursalID", "Saldo"}

	for i, cajero := range cajeros {

		resultados[i] = []string{
			fmt.Sprintf("%d", cajero.ID),
			ObtenerUsuarioPorCajeroId(*cajero.UsuarioID),
			ObtenerNombreSucursalPorLaId(cajero.SucursalID),
			FormatearDinero(cajero.Saldo),
		}
	}

	return resultados, nil
}

// Enlistar cajeros por sucursal id
func EnlistarCajerosPorSucursalId(sucursal_id int) ([][]string, error) {
	// Verificar que la sucursal_id no sea 0 o menor a 0
	if sucursal_id <= 0 {
		return nil, errors.New("La sucursal_id no es válida")
	}
	// Verificar que la sucursal exista
	if !SucursalExistePorLaId(sucursal_id) {
		return nil, errors.New("La sucursal no existe")
	}

	var cajeros []Cajero
	if err := Db.Where("sucursal_id = ?", sucursal_id).Find(&cajeros).Error; err != nil {
		println(err.Error())
		return nil, err
	}

	resultados := make([][]string, len(cajeros))

	for i, cajero := range cajeros {
		resultados[i] = []string{
			fmt.Sprintf("%d", cajero.ID),
			ObtenerUsuarioPorCajeroId(cajero.ID),
			ObtenerNombreSucursalPorLaId(cajero.SucursalID),
			FormatearDinero(cajero.Saldo),
		}
	}
	return resultados, nil
}

// Obtener el nombre del usuario actual en el cajero teniendo en cuenta la id del cajero, devolver el string "DESOCUPADO" si no existe
func ObtenerUsuarioPorCajeroId(id int) string {
	var cajero Cajero
	Db.Table("cajeros").Where("id = ?", id).First(&cajero)

	var usuario Usuarios
	result := Db.Table("usuarios").Where("id = ?", cajero.UsuarioID).First(&usuario)
	if result.Error != nil {
		return "DESOCUPADO"
	} else {
		return usuario.Nombre
	}

}

// CajeroExistePorLaId
func CajeroExistePorLaId(id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var s Cajero
	result := Db.Table("cajeros").Where("id = ?", id).First(&s)
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
	var s Cajero
	result := Db.Table("cajeros").Where("usuario_id = ?", usuario_id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

// Verificar que el usuario este en el cajero
func UsuarioEstaEnElCajero(usuario_id, cajero_id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if usuario_id <= 0 || cajero_id <= 0 {
		return false
	}
	// Verificar que el usuario existe
	if !UsuarioExistePorLaId(usuario_id) {
		return false
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(cajero_id) {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var s Cajero
	result := Db.Table("cajeros").Where("usuario_id = ? AND id = ?", usuario_id, cajero_id).First(&s)
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

	var s Cajero
	result := Db.Table("cajeros").Where("id = ?", id).First(&s)
	if result.Error != nil {
		return 0, result.Error
	}
	return s.Saldo, nil
}

func CajeroPerteneceASucursal(cajero_id, sucursal_id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if cajero_id <= 0 || sucursal_id <= 0 {
		return false
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(cajero_id) {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var s Cajero
	result := Db.Table("cajeros").Where("id = ? AND sucursal_id = ?", cajero_id, sucursal_id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

/* Movimientos caja */

// CrearSesion
func CrearSesion(usuario_id, cajero_id, sucursal_id, clave, entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50 int) (string, error) {
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(usuario_id) {
		return "", errors.New("El usuario no existe")
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(cajero_id) {
		return "", errors.New("El cajero no existe")
	}
	// Verificar que el usuario no este en un cajero distinto
	if UsuarioEnSesion(usuario_id) {
		return "", errors.New("El usuario ya esta en sesion")
	}
	// Verificar que el cajero pertenezca a la sucursal
	if !CajeroPerteneceASucursal(cajero_id, sucursal_id) {
		return "", errors.New("El cajero no pertenece a la sucursal")
	}
	// Verificar que la sucursal exista
	if !SucursalExistePorLaId(sucursal_id) {
		return "", errors.New("La sucursal no existe")
	}
	// Verificar que la clave sea correcta
	if !ClaveCorrecta(usuario_id, clave) {
		return "", errors.New("La clave es incorrecta")
	}
	// Verificar que el saldo de inicio sea el mismo saldo que el cajero cerro la sesion anterior
	// TODO: Hacer esto

	// TODO: Pasar a gorm
	result := Db.Exec("INSERT INTO movimientos_caja (cajero_id, usuario_id, fecha_entrada, entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", cajero_id, usuario_id, time.Now(), entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50)
	if result.Error != nil {
		return "", result.Error
	}

	// Asignar usuario a cajero
	result = Db.Table("cajeros").Where("id = ?", cajero_id).Update("usuario_id", usuario_id)
	if result.Error != nil {
		return "", result.Error
	}
	// Actualizar usuario
	result = Db.Table("usuarios").Where("id = ?", usuario_id).Update("usuario_en_sesion", true)
	if result.Error != nil {
		return "", result.Error
	}

	saldos, err := SumarSaldos(entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50)
	if err != nil {
		return "", err
	}

	// Actualizar saldo en cajero
	result = Db.Table("cajeros").Where("id = ?", cajero_id).Update("saldo", saldos)
	if result.Error != nil {
		return "", result.Error
	}
	statusMessage := fmt.Sprintf("Sesion creada con un saldo de %d pesos", saldos)
	return statusMessage, nil
}

// CerrarSesion
func CerrarSesion(usuario_id, cajero_id, sucursal_id, salida_billetes_100, salida_billetes_50, salida_billetes_20, salida_billetes_10, salida_billetes_5, salida_billetes_2, salida_billetes_1, salida_monedas_1000, salida_monedas_500, salida_monedas_200, salida_monedas_100, salida_monedas_50 int) (string, error) {
	var statusMessage string
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(usuario_id) {
		return "", errors.New("El usuario no existe")
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(cajero_id) {
		return "", errors.New("El cajero no existe")
	}
	// Verificar que el usuario este en el cajero
	if !UsuarioEstaEnElCajero(cajero_id, usuario_id) {
		return "", errors.New("El usuario no esta en el cajero")
	}
	// Verificar que la sucursal exista
	if !SucursalExistePorLaId(sucursal_id) {
		return "", errors.New("La sucursal no existe")
	}

	saldosSalida, err := SumarSaldos(salida_billetes_100, salida_billetes_50, salida_billetes_20, salida_billetes_10, salida_billetes_5, salida_billetes_2, salida_billetes_1, salida_monedas_1000, salida_monedas_500, salida_monedas_200, salida_monedas_100, salida_monedas_50)
	if err != nil {
		return "", err
	}
	// Verificar que el saldo de salida sea el mismo que esta en el cajero
	descuadre, resultSaldo := SaldoEsIgualAlSaldoDelCajero(saldosSalida, cajero_id)

	// Actualizar usuario quitandole el ensesion
	result := Db.Table("usuarios").Where("id = ?", usuario_id).Update("usuario_en_sesion", false)
	if result.Error != nil {
		return "", result.Error
	}
	// Actualizar cajero quitandole el usuario
	result = Db.Table("cajeros").Where("id = ?", cajero_id).Update("usuario_id", nil)
	if result.Error != nil {
		return "", result.Error
	}
	// Actualizar el movimiento y agregarle los salidas
	result = Db.Exec("UPDATE movimientos_caja SET fecha_salida = ?, salida_billetes_100 = ?, salida_billetes_50 = ?, salida_billetes_20 = ?, salida_billetes_10 = ?, salida_billetes_5 = ?, salida_billetes_2 = ?, salida_billetes_1 = ?, salida_monedas_1000 = ?, salida_monedas_500 = ?, salida_monedas_200 = ?, salida_monedas_100 = ?, salida_monedas_50 = ? WHERE cajero_id = ? AND usuario_id = ? AND fecha_salida IS NULL", time.Now(), salida_billetes_100, salida_billetes_50, salida_billetes_20, salida_billetes_10, salida_billetes_5, salida_billetes_2, salida_billetes_1, salida_monedas_1000, salida_monedas_500, salida_monedas_200, salida_monedas_100, salida_monedas_50, cajero_id, usuario_id)
	if result.Error != nil {
		return "", result.Error
	}

	if !resultSaldo {
		statusMessage = fmt.Sprintf("El saldo de salida es diferente al saldo del cajero, el descuadre es de %d pesos", descuadre)
	} else {
		statusMessage = fmt.Sprintf("Sesion cerrada con un saldo de %d pesos", saldosSalida)
	}
	return statusMessage, nil
}

func SaldoEsIgualAlSaldoDelCajero(saldoSalida, cajero_id int) (*int, bool) {
	// Verificar que el saldo no sea menor a 0
	if saldoSalida < 0 {
		return nil, false
	}
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(cajero_id) {
		return nil, false
	}

	// Verificar que el saldo sea igual al saldo del cajero
	var c Cajero
	result := Db.Table("cajeros").Where("id = ?", cajero_id).First(&c)
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

/* Plataformas */

// CrearPlataformas
func CrearPlataformas(nombre string, sucursal_id int) (string, error) {
	// Verificar que el nombre no este vacio
	if nombre == "" {
		return "", errors.New("El nombre no puede estar vacio")
	}
	// Verificar que la sucursal no sea 0 o menor a 0
	if sucursal_id <= 0 {
		return "", errors.New("La sucursal no puede ser 0 o menor a 0")
	}
	// Verificar que la sucursal exista
	if SucursalExistePorLaId(sucursal_id) == false {
		return "", errors.New("La sucursal no existe")
	}

	// Crear la plataforma
	result := Db.Table("plataformas").Create(&Plataformas{
		SucursalID:    sucursal_id,
		Nombre:        nombre,
		FechaCreacion: time.Now(),
	})
	if result.Error != nil {
		return "", result.Error
	}
	return "Plataforma creada exitosamente", nil
}

// EliminarPlataforma
func EliminarPlataforma(id int) (string, error) {
	// Verificar que el id no sea 0 o menor a 0
	if id <= 0 {
		return "", errors.New("El id no puede ser 0 o menor a 0")
	}
	// Verificar que la plataforma exista
	if PlataformaExistePorLaId(id) == false {
		return "", errors.New("La plataforma no existe")
	}
	// Eliminar la plataforma
	result := Db.Table("plataformas").Where("id = ?", id).Delete(&Plataformas{})
	if result.Error != nil {
		return "", result.Error
	}
	return "Plataforma eliminada exitosamente", nil
}

// EditarPlataforma
func EditarPlataforma(id_plataforma int, tipo string, valor any) (string, error) {
	// Verificar que la plataforma sea no sea 0 o menor a 0
	if id_plataforma <= 0 {
		return "", errors.New("El id de la plataforma no puede ser 0 o menor a 0")
	}
	// Verificar que la platafora exista
	if PlataformaExistePorLaId(id_plataforma) == false {
		return "", errors.New("La plataforma no existe")
	}
	// Verificar que el tipo sea valido
	TiposValidos := []string{"nombre", "sucursal_id"}
	if !Contains(TiposValidos, tipo) {
		return "", errors.New("El tipo no es valido")
	}

	// Editar la plataforma
	result := Db.Table("plataformas").Where("id = ?", id_plataforma).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Plataforma editada exitosamente", nil
}

// EnlistarPlataformas
func EnlistarPlataformas() ([]Plataformas, error) {
	var plataformas []Plataformas
	result := Db.Table("plataformas").Find(&plataformas)
	if result.Error != nil {
		return nil, result.Error
	}
	return plataformas, nil
}

// PlataformaExistePorLaId
func PlataformaExistePorLaId(id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var p Plataformas
	result := Db.Table("plataformas").Where("id = ?", id).First(&p)
	if result.Error != nil {
		return false
	}
	return true
}

/* Sucursales */

func CrearSucursal(nombre string) (string, error) {
	// Verificar que el nombre no sea vacio
	if nombre == "" {
		return "", errors.New("El nombre no puede ser vacio")
	}
	// Verificar que el nombre no exista
	if SucursalExistePorElNombre(nombre) {
		return "", errors.New("Ya existe una sucursal con este nombre")
	}

	// Crear la sucursal
	result := Db.Table("sucursales").Create(&Sucursal{Nombre: nombre, FechaCreacion: time.Now()})
	if result.Error != nil {
		return "", result.Error
	}
	// Obtener los datos de la nueva sucursal
	var sucursal Sucursal
	result = Db.Table("sucursales").Where("nombre = ?", nombre).First(&sucursal)
	if result.Error != nil {
		return "", result.Error
	}
	// Crear la caja menor de la sucursal
	result = Db.Table("caja_menor").Create(&CajaMenor{
		SucursalID: sucursal.ID,
		Saldo:      0})
	if result.Error != nil {
		return "", result.Error
	}
	return "Sucursal creada exitosamente", nil
}

func EliminarSucursal(nombre string) (string, error) {
	// Verificar que el nombre no sea vacio
	if nombre == "" {
		return "", errors.New("El nombre no puede ser vacio")
	}
	// Verificar que el nombre exista
	if !SucursalExistePorElNombre(nombre) {
		return "", errors.New("No existe una sucursal con este nombre")
	}

	// Obtener el id de la sucursal por el nombre
	var sucursal Sucursal
	result := Db.Table("sucursales").Where("nombre = ?", nombre).First(&sucursal)
	if result.Error != nil {
		return "", result.Error
	}
	result = Db.Table("caja_menor").Where("sucursal_id = ?", sucursal.ID).Delete(&CajaMenor{})
	// Eliminar la sucursal
	result = Db.Table("sucursales").Where("nombre = ?", nombre).Delete(&Sucursal{})
	if result.Error != nil {
		return "", result.Error
	}
	// Eliminar caja menor de la sucursal

	return "Sucursal eliminada exitosamente", nil
}

func EditarSucursal(id int, tipo, valor string) (string, error) {
	// Verificar que el tipo y el valor no esten vacios
	if tipo == "" || valor == "" {
		return "", errors.New("El tipo y el valor no pueden ser vacios")
	}
	// Verificar que la id no sea 0 o menor a 0
	if id <= 0 {
		return "", errors.New("La id no puede ser 0 o menor a 0")
	}
	// Verificar que el id exista
	if !SucursalExistePorLaId(id) {
		return "", errors.New("No existe una sucursal con este id")
	}
	if tipo == "nombre" {
		// Verificar que el nombre no exista
		if SucursalExistePorElNombre(valor) {
			return "", errors.New("Ya existe una sucursal con este nombre")
		}
	}
	// Verificar que el tipo sea valido dentro de un array de tipos de la variable TiposSucursales
	TiposSucursales := []string{"nombre"}
	if !Contains(TiposSucursales, tipo) {
		return "", errors.New("El tipo no es valido")
	}

	// Editar la sucursal
	result := Db.Table("sucursales").Where("id = ?", id).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Sucursal editada exitosamente", nil
}

func EnlistarSucursales() ([][]string, error) {
	var sucursales []Sucursal
	result := Db.Table("sucursales").Find(&sucursales)
	if result.Error != nil {
		return nil, result.Error
	}

	var sucursalesEnlistadas [][]string
	for _, sucursal := range sucursales {
		sucursalesEnlistadas = append(sucursalesEnlistadas, []string{
			strconv.Itoa(sucursal.ID),
			sucursal.Nombre,
			sucursal.FechaCreacion.Format("2006-01-02 15:04:05")})
	}
	return sucursalesEnlistadas, nil
}

// Verificar que la sucursal no exista
func SucursalExistePorElNombre(nombre string) bool {
	// Verificar que el nombre no sea vacio
	if nombre == "" {
		return false
	}
	// Verificar que el nombre no exista en la base de datos
	var s Sucursal
	result := Db.Table("sucursales").Where("nombre = ?", nombre).First(&s)
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
	var s Sucursal
	result := Db.Table("sucursales").Where("id = ?", id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

// Obtener el nombre de la sucursal por la id
func ObtenerNombreSucursalPorLaId(id int) string {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return ""
	}
	// Verificar que el id exista en la base de datos
	var s Sucursal
	result := Db.Table("sucursales").Where("id = ?", id).First(&s)
	if result.Error != nil {
		return ""
	}
	return s.Nombre
}

/* Transacciones */

// CrearTransacciones
func CrearTransaccion(cajero_id, usuario_id, valor int, tipo, descripcion string, billetera bool) (string, error) {
	// Verificar que el cajero exista
	if !CajeroExistePorLaId(cajero_id) {
		return "El cajero no existe", nil
	}
	// Verificar que la billetera sea false o true
	if billetera != false && billetera != true {
		return "La billetera no es valida", nil
	}
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(usuario_id) {
		return "El usuario no existe", nil
	}
	// Verificar que el usuario este en el cajero
	if !UsuarioEstaEnElCajero(usuario_id, cajero_id) {
		return "El usuario no esta en el cajero", nil
	}
	// Verificar que el valor no sea menor a 0
	if valor <= 0 {
		return "El valor no puede ser menor a 0", nil
	}
	// Verificar que el tipo no sea vacio
	Tipos := []string{"DEPOSITO", "RETIRO"}
	if !Contains(Tipos, tipo) {
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
		saldo, err := DevolverSaldoCajeroPorLaId(cajero_id)
		if err != nil {
			return "", err
		}
		// Sumar el valor al saldo del cajero
		saldo += valor
		// Actualizar el saldo del cajero
		result := Db.Table("cajeros").Where("id = ?", cajero_id).Update("saldo", saldo)
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
		saldo, err := DevolverSaldoCajeroPorLaId(cajero_id)
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
		result := Db.Table("cajeros").Where("id = ?", cajero_id).Update("saldo", saldo)
		if result.Error != nil {
			return "", result.Error
		}
	}
	// Crear la transaccion
	result := Db.Table("transacciones").Create(&Transacciones{
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
	result := Db.Table("transacciones").Where("id = ?", id).Delete(&Transacciones{})
	if result.Error != nil {
		return "", result.Error
	}
	return "Transaccion eliminada", nil
}

// EditarTransaccion
func EditarTransaccion(usuario_id, transaccion_id int, tipo string, valor any) (string, error) {
	// Verificar que el usuario tenga el rango adecuado
	if !RangoCorrecto(usuario_id) {
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
	if !Contains(Tipos, tipo) {
		return "El tipo no es valido", nil
	}
	if tipo == "descripcion" {
		// Verificar que la descripcion no sea vacia
		if valor == "" {
			descripcion := "Sin descripcion"
			valor = descripcion
		} else {
			descripcion := valor.(string)
			valor = descripcion
		}
		// Actualizar la descripcion de la transaccion
		result := Db.Table("transacciones").Where("id = ?", transaccion_id).Update("descripcion", valor)
		if result.Error != nil {
			return "", result.Error
		}
	} else if tipo == "valor" {
		// Verificar que el valor no sea vacio o menor a 0
		if valor.(int) <= 0 {
			return "", errors.New("El valor no puede ser vacio o menor a 0")
		}
		// Actualizar el valor de la transaccion
		result := Db.Table("transacciones").Where("id = ?", transaccion_id).Update("valor", valor)
		if result.Error != nil {
			return "", result.Error
		}
	}
	// Modificar fecha_actualizacion
	result := Db.Table("transacciones").Where("id = ?", transaccion_id).Update("fecha_actualizacion", time.Now())
	if result.Error != nil {
		return "", result.Error
	}
	return "Transaccion editada", nil
}

// EnlistarTransacciones devuelve todas las transacciones en un [][]string
func EnlistarTransacciones() ([][]string, error) {
	// Obtener todas las transacciones de la tabla
	var transacciones []Transacciones
	if err := Db.Find(&transacciones).Error; err != nil {
		return nil, err
	}

	// Construir la matriz de resultados
	resultados := make([][]string, len(transacciones)+1)
	resultados[0] = []string{"Cajero", "Usuario", "Tipo", "Descripcion", "Billetera", "Valor", "Fecha creacion", "Fecha actualizacion", "Editar", "Eliminar"}

	for i, transaccion := range transacciones {
		resultados[i+1] = []string{
			fmt.Sprintf("%d", transaccion.CajeroID),
			NombrePorLaId(transaccion.UsuarioID),
			transaccion.Tipo,
			transaccion.Descripcion,
			FormatearBilletera(transaccion.Billetera),
			FormatearDinero(transaccion.Valor),
			FormatearFecha(transaccion.FechaCreacion),
			FormatearFecha(transaccion.FechaActualizacion),
			"Editar",
			"Eliminar",
		}
	}

	return resultados, nil
}

func ElistarTransaccionesPorSucursal(sucursal_id int) ([][]string, error) {
	// Obtener el id de los cajeros de la sucursal
	var cajeros []Cajero
	if err := Db.Table("cajeros").Where("sucursal_id = ?", sucursal_id).Find(&cajeros).Error; err != nil {
		return nil, err
	}
	// Obtener los ID de los cajeros
	var cajeroIDs []int
	for _, cajero := range cajeros {
		cajeroIDs = append(cajeroIDs, cajero.ID)
	}
	// Obtener todas las transacciones de los cajeros de la sucursal
	var transacciones []Transacciones
	if err := Db.Table("transacciones").Where("cajero_id IN (?)", cajeroIDs).Find(&transacciones).Error; err != nil {
		return nil, err
	}
	// Construir la matriz de resultados
	resultados := make([][]string, len(transacciones)+1)
	resultados[0] = []string{"ID", "Cajero", "Usuario", "Tipo", "Descripcion", "Billetera", "Valor", "Fecha creacion", "Fecha actualizacion", "Editar", "Eliminar"}

	for i, transaccion := range transacciones {
		resultados[i+1] = []string{
			fmt.Sprintf("%d", transaccion.ID),
			fmt.Sprintf("%d", transaccion.CajeroID),
			NombrePorLaId(transaccion.UsuarioID),
			transaccion.Tipo,
			transaccion.Descripcion,
			FormatearBilletera(transaccion.Billetera),
			FormatearDinero(transaccion.Valor),
			FormatearFecha(transaccion.FechaCreacion),
			FormatearFecha(transaccion.FechaActualizacion),
			"Editar",
			"Eliminar",
		}
	}

	return resultados, nil
}

func FormatearBilletera(billetera bool) string {
	if billetera {
		return "Si"
	}
	return "No"
}

// Si la fecha esta vacia, di que no hay fecha
func FormatearFecha(fecha time.Time) string {
	if fecha.IsZero() {
		return "No se ha actualizado"
	}
	return fecha.Format("02/01/2006 - 15:04")
}

// Verificar que la transaccion exista
func TransaccionExistePorLaId(transaccion_id int) bool {
	// Verificar que la id no sea 0 o menor a 0
	if transaccion_id <= 0 {
		return false
	}
	// Verificar que la transaccion exista
	var s Transacciones
	result := Db.Table("transacciones").Where("id = ?", transaccion_id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}

/* Usuarios */

// Crear usuarios
func CrearUsuario(usuario, nombre string, clave, rango int) (string, error) {
	// Verificar que usuario y nombre no este vacio
	if usuario == "" || nombre == "" {
		return "", errors.New("El usuario y el nombre no pueden estar vacios")
	}
	// Verificar que la clave y el rango sean validas, el rango solo esta el rango 1 y 2 y la clave tiene que ser de 4 digitos
	if clave < 1000 || clave > 9999 || rango < 1 || rango > 2 {
		return "", errors.New("La clave y el rango no son validos")
	}
	// Verificar que el usuario no exista
	if UsuarioExistePorElUsuario(usuario) {
		return "", errors.New("El usuario ya existe")
	}

	// Crear el usuario
	result := Db.Table("usuarios").Create(&Usuarios{
		Usuario:         usuario,
		Nombre:          nombre,
		Clave:           clave,
		Rango:           rango,
		UsuarioEnSesion: false,
		FechaCreacion:   time.Now(),
	})
	if result.Error != nil {
		return "", result.Error
	}
	return "Usuario creado", nil
}

// Eliminar usuarios
func EliminarUsuario(usuario string) (string, error) {
	// Verificar que el usuario no este vacio
	if usuario == "" {
		return "", errors.New("El usuario no puede estar vacio")
	}
	// Verificar que el usuario exista
	if !UsuarioExistePorElUsuario(usuario) {
		return "", errors.New("El usuario no existe")
	}

	// Eliminar el usuario
	result := Db.Table("usuarios").Where("usuario = ?", usuario).Delete(&Usuarios{})
	if result.Error != nil {
		return "", result.Error
	}
	return "Usuario eliminado", nil
}

// Editar usuario
func EditarUsuario(id int, tipo string, valor any) (string, error) {
	// Verificar que la id no sea 0 o menor a 0
	if id <= 0 {
		return "", errors.New("La id no es valida")
	}
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(id) {
		return "", errors.New("El usuario no existe")
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
	TiposUsuarios := []string{"usuario", "nombre", "clave", "rango"}
	if !Contains(TiposUsuarios, tipo) {
		return "", errors.New("El tipo no es valido")
	}

	// Editar el usuario
	result := Db.Table("usuarios").Where("id = ?", id).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Usuario editado", nil
}

// Enlistar usuarios
func EnlistarUsuarios() ([][]string, error) {
	// Obtener todas las cajeros de la tabla
	var usuarios []Usuarios
	if err := Db.Find(&usuarios).Error; err != nil {
		return nil, err
	}

	// Construir la matriz de resultados
	resultados := make([][]string, len(usuarios)+1)
	resultados[0] = []string{"ID", "Usuario", "Nombre", "Clave", "Rango", "Fecha creacion", "Editar", "Eliminar"}

	for i, usuario := range usuarios {
		resultados[i+1] = []string{
			fmt.Sprintf("%d", usuario.Id),
			usuario.Usuario,
			usuario.Nombre,
			fmt.Sprintf("%d", usuario.Clave),
			fmt.Sprintf("%d", usuario.Rango),
			FormatearFecha(usuario.FechaCreacion),
			"Editar",
			"Eliminar",
		}
	}

	return resultados, nil
}

func UsuarioExistePorElUsuario(usuario string) bool {
	// Verificar que el usuario no este vacio
	if usuario == "" {
		return false
	}
	// Verificar que el usuario no exista
	var u Usuarios
	result := Db.Table("usuarios").Where("usuario = ?", usuario).First(&u)
	if result.Error != nil {
		return false
	}
	return true
}

func UsuarioExistePorLaId(id int) bool {
	// Verificar que la id no sea 0 o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el usuario no exista
	var u Usuarios
	result := Db.Table("usuarios").Where("id = ?", id).First(&u)
	if result.Error != nil {
		return false
	}
	return true
}

func UsuarioEnSesion(id int) bool {
	// Verificar que la id no sea 0 o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el usuario no exista
	var u Usuarios
	result := Db.Table("usuarios").Where("id = ?", id).First(&u)
	if result.Error != nil {
		return false
	}
	return u.UsuarioEnSesion
}

// Verificar que la clave sea correcta
func ClaveCorrecta(usuario_id, clave int) bool {
	// Verificar que la id no sea 0 o menor a 0
	if usuario_id <= 0 {
		return false
	}
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(usuario_id) {
		return false
	}
	// Verificar que la clave sea correcta
	var u Usuarios
	result := Db.Table("usuarios").Where("id = ?", usuario_id).First(&u)
	if result.Error != nil {
		return false
	}
	if u.Clave != clave {
		return false
	}
	return true
}

// Verificar que el usuario tenga el rango correcto
func RangoCorrecto(usuario_id int) bool {
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(usuario_id) {
		return false
	}
	// Verificar que el usuario tenga el rango correcto
	var u Usuarios
	result := Db.Table("usuarios").Where("id = ?", usuario_id).First(&u)
	if result.Error != nil {
		return false
	}
	// Si el rango es igual o mayor a 2 significa que si tiene los permisos adecuados
	fmt.Println(u.Rango)
	if u.Rango >= 2 {
		return true
	}
	return false
}

// Devolver el nombre del usuario por la id
func NombrePorLaId(usuario_id int) string {
	// Si el usuario no existe devolver un string vacio
	if !UsuarioExistePorLaId(usuario_id) {
		return "Usuario no encontrado"
	}
	// Devolver el nombre del usuario
	var u Usuarios
	result := Db.Table("usuarios").Where("id = ?", usuario_id).First(&u)
	if result.Error != nil {
		return "Usuario no encontrado"
	}
	return u.Nombre
}
