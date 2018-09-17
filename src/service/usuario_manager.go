package service

import (
	"errors"

	"github.com/francoagarcia/meli-go/src/domain"
)

// UsuarioManager tweet manager
type UsuarioManager struct {
	Usuarios []*domain.Usuario
}

// NewUsuarioManager inicializar slices
func NewUsuarioManager() *UsuarioManager {
	usuarioManager := UsuarioManager{Usuarios: make([]*domain.Usuario, 0)}
	return &usuarioManager
}

// RegistrarUsuario nuevo usuario
func (u *UsuarioManager) RegistrarUsuario(usuario *domain.Usuario) error {
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
	u.Usuarios = append(u.Usuarios, usuario)
	return nil
}

// GetUsuarioByUsername buscar usuario por username
func (u *UsuarioManager) GetUsuarioByUsername(username string) *domain.Usuario {
	for _, aUser := range u.Usuarios {
		if aUser.Username == username {
			return aUser
		}
	}
	return nil
}

// IsUserRegistered validar si usuario está registrado por username
//func IsUserRegistered(username string) bool {
//	usuarioBuscado := GetUsuarioByUsername(username)
//	return usuarioBuscado != nil
//}
