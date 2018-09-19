package domain

import (
	"github.com/satori/go.uuid"
)

// Usuario struct
type Usuario struct {
	Username    string     `json:"username"`
	Nombre      string     `json:"nombre"`
	Mail        string     `json:"mail"`
	Contrasenia string     `json:"contrasenia"`
	ID          *uuid.UUID `json:"id"`
}

// NewUsuario nuevo tweet
func NewUsuario(username, nombre, mail, contrasenia string) *Usuario {
	uuid, _ := uuid.NewV4()
	var usuario = Usuario{
		Username:    username,
		Nombre:      nombre,
		Mail:        mail,
		Contrasenia: contrasenia,
		ID:          &uuid}
	return &usuario
}
