package plataformas

import (
	"errors"
	"julio/db"
	"julio/models"
	"time"
)

// AsignarSaldoPlataforma
func AsignarSaldoPlataforma(id_plataforma int, saldo int) (string, error) {
	// Verificar que la plataforma no sea 0 o menor a 0
	if id_plataforma <= 0 {
		return "", errors.New("El id no puede ser 0 o menor a 0")
	}
	// Verificar que la plataforma exista
	if PlataformaExistePorLaId(id_plataforma) == false {
		return "", errors.New("La plataforma no existe")
	}
	// Verificar que el saldo no sea 0 o menor a 0
	if saldo == 0 {
		return "", errors.New("El saldo no puede ser 0 o menor a 0")
	}

	// Asignar el saldo a la plataforma
	result := db.Db.Table("saldos_plataforma").Create(&models.SaldosPlataforma{
		PlataformaID:  id_plataforma,
		Saldo:         saldo,
		FechaCreacion: time.Now(),
	})
	if result.Error != nil {
		return "", result.Error
	}
	return "Saldo asignado exitosamente", nil
}
