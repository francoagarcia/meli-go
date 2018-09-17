package service

import (
	"errors"

	"github.com/francoagarcia/meli-go/src/domain"
)

// Usuarios registrados
var Usuarios []*domain.Usuario

// RegistrarUsuario nuevo usuario
func RegistrarUsuario(usuario *domain.Usuario) error {
	if usuario.Nombre == "" {
		return errors.New("nombre is required")
	}
	if usuario.Mail == "" {
		return errors.New("mail is required")
	}
	if usuario.Username == "" {
		return errors.New("username is required")
	}
	if usuario.Contrasenia == "" {
		return errors.New("contrasenia is required")
	}
	Usuarios = append(Usuarios, usuario)
	return nil
}

// GetUsuarioByUsername buscar usuario por username
func GetUsuarioByUsername(username string) *domain.Usuario {
	for i := 0; i < len(Usuarios); i++ {
		if Usuarios[i].Username == username {
			return Usuarios[i]
		}
	}
	return nil
}
