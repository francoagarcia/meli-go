package domain

// Usuario struct
type Usuario struct {
	Username    string
	Nombre      string
	Mail        string
	Contrasenia string
}

// NewUsuario nuevo tweet
func NewUsuario(username, nombre, mail, contrasenia string) *Usuario {
	var usuario = Usuario{
		username,
		nombre,
		mail,
		contrasenia}
	return &usuario
}
