package usuarios

import (
	"julio/db"
	"julio/models"
)

func UsuarioExistePorElUsuario(usuario string) bool {
	// Verificar que el usuario no este vacio
	if usuario == "" {
		return false
	}
	// Verificar que el usuario no exista
	var u models.Usuarios
	result := db.Db.Table("usuarios").Where("usuario = ?", usuario).First(&u)
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
	var u models.Usuarios
	result := db.Db.Table("usuarios").Where("id = ?", id).First(&u)
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
	var u models.Usuarios
	result := db.Db.Table("usuarios").Where("id = ?", id).First(&u)
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
	var u models.Usuarios
	result := db.Db.Table("usuarios").Where("id = ?", usuario_id).First(&u)
	if result.Error != nil {
		return false
	}
	if u.Clave != clave {
		return false
	}
	return true
}
