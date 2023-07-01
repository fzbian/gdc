package plataformas

import (
	"errors"
	"julio/controllers/sucursales"
	"julio/db"
	"julio/models"
	"julio/utils"
	"time"
)

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
	if sucursales.SucursalExistePorLaId(sucursal_id) == false {
		return "", errors.New("La sucursal no existe")
	}

	// Crear la plataforma
	result := db.Db.Table("plataformas").Create(&models.Plataformas{
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
	result := db.Db.Table("plataformas").Where("id = ?", id).Delete(&models.Plataformas{})
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
	if !utils.Contains(TiposValidos, tipo) {
		return "", errors.New("El tipo no es valido")
	}

	// Editar la plataforma
	result := db.Db.Table("plataformas").Where("id = ?", id_plataforma).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Plataforma editada exitosamente", nil
}

// EnlistarPlataformas
func EnlistarPlataformas() ([]models.Plataformas, error) {
	var plataformas []models.Plataformas
	result := db.Db.Table("plataformas").Find(&plataformas)
	if result.Error != nil {
		return nil, result.Error
	}
	return plataformas, nil
}
