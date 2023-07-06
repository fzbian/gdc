package usuarios

import (
	"fmt"
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

// Verificar que el usuario tenga el rango correcto
func RangoCorrecto(usuario_id int) bool {
	// Verificar que el usuario exista
	if !UsuarioExistePorLaId(usuario_id) {
		return false
	}
	// Verificar que el usuario tenga el rango correcto
	var u models.Usuarios
	result := db.Db.Table("usuarios").Where("id = ?", usuario_id).First(&u)
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
	var u models.Usuarios
	result := db.Db.Table("usuarios").Where("id = ?", usuario_id).First(&u)
	if result.Error != nil {
		return "Usuario no encontrado"
	}
	return u.Nombre
}
