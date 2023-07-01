package plataformas

import (
	"julio/db"
	"julio/models"
)

// PlataformaExistePorLaId
func PlataformaExistePorLaId(id int) bool {
	// Verificar que el id no sea vacio o menor a 0
	if id <= 0 {
		return false
	}
	// Verificar que el id no exista en la base de datos
	var p models.Plataformas
	result := db.Db.Table("plataformas").Where("id = ?", id).First(&p)
	if result.Error != nil {
		return false
	}
	return true
}
