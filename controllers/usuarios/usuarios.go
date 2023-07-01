package usuarios

import (
	"errors"
	"julio/db"
	"julio/models"
	"julio/utils"
	"time"
)

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
	result := db.Db.Table("usuarios").Create(&models.Usuarios{
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
	result := db.Db.Table("usuarios").Where("usuario = ?", usuario).Delete(&models.Usuarios{})
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
	if !utils.Contains(TiposUsuarios, tipo) {
		return "", errors.New("El tipo no es valido")
	}

	// Editar el usuario
	result := db.Db.Table("usuarios").Where("id = ?", id).Update(tipo, valor)
	if result.Error != nil {
		return "", result.Error
	}
	return "Usuario editado", nil
}

// Enlistar usuarios
func EnlistarUsuarios() ([]models.Usuarios, error) {
	var usuarios []models.Usuarios
	result := db.Db.Table("usuarios").Find(&usuarios)
	if result.Error != nil {
		return nil, result.Error
	}
	return usuarios, nil
}
