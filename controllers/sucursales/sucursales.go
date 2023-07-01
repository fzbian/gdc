package sucursales

import (
	"errors"
	"julio/db"
	"julio/models"
	"julio/utils"
	"time"
)

func CreateSucursal(nombre string) (string, error) {
	// Verificar que el nombre no sea vacio
	if nombre == "" {
		return "", errors.New("El nombre no puede ser vacio")
	}
	// Verificar que el nombre no exista
	if SucursalExistePorElNombre(nombre) {
		return "", errors.New("Ya existe una sucursal con este nombre")
	}

	// Crear la sucursal
	result := db.Db.Table("sucursales").Create(&models.Sucursal{Nombre: nombre, FechaCreacion: time.Now()})
	if result.Error != nil {
		return "", result.Error
	}
	// Obtener los datos de la nueva sucursal
	var sucursal models.Sucursal
	result = db.Db.Table("sucursales").Where("nombre = ?", nombre).First(&sucursal)
	if result.Error != nil {
		return "", result.Error
	}
	// Crear la caja menor de la sucursal
	result = db.Db.Table("caja_menor").Create(&models.CajaMenor{
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

	// Eliminar la sucursal
	result := db.Db.Table("sucursales").Where("nombre = ?", nombre).Delete(&models.Sucursal{})
	if result.Error != nil {
		return "", result.Error
	}
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
	// Verificar que el tipo sea valido dentro de un array de tipos de la variable TiposSucursales
	TiposSucursales := []string{"nombre"}
	if !utils.Contains(TiposSucursales, tipo) {
		return "", errors.New("El tipo no es valido")
	}

	// Editar la sucursal
	result := db.Db.Table("sucursales").Where("id = ?", id).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Sucursal editada exitosamente", nil
}

func EnlistarSucursales() ([]models.Sucursal, error) {
	var sucursales []models.Sucursal
	result := db.Db.Table("sucursales").Find(&sucursales)
	if result.Error != nil {
		return nil, result.Error
	}
	return sucursales, nil
}
