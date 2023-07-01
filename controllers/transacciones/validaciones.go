package transacciones

import (
	"julio/db"
	"julio/models"
)

// Verificar que la transaccion exista
func TransaccionExistePorLaId(transaccion_id int) bool {
	// Verificar que la id no sea 0 o menor a 0
	if transaccion_id <= 0 {
		return false
	}
	// Verificar que la transaccion exista
	var s models.Transacciones
	result := db.Db.Table("transacciones").Where("id = ?", transaccion_id).First(&s)
	if result.Error != nil {
		return false
	}
	return true
}
