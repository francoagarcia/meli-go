package domain

// Usuario struct
type Usuario struct {
	Username    string `json:"username"`
	Nombre      string `json:"nombre"`
	Mail        string `json:"mail"`
	Contrasenia string `json:"contrasenia"`
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
