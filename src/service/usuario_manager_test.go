package service_test

import (
	"testing"

	"github.com/francoagarcia/meli-go/src/domain"
	"github.com/francoagarcia/meli-go/src/service"
)

func TestNoSeEncuentraUsuarioPorID(t *testing.T) {
	// Initialization
	usuarioManager := service.NewUsuarioManager()

	// Operation
	usuarioRegistrado := usuarioManager.GetUsuarioByUsername("usernameNoExistente")

	// Validation
	if usuarioRegistrado != nil {
		t.Error("Expected usuario must be nil")
	}
}
func TestRegistrarUsuario(t *testing.T) {

	// Initialization
	usuarioManager := service.NewUsuarioManager()
	var usuario *domain.Usuario

	var nombre string = "Franco"
	var username string = "francoagarcia"
	var mail string = "francoagarcia@gmail.com"
	var contrasenia string = "pass123"

	usuario = domain.NewUsuario(username, nombre, mail, contrasenia)

	// Operation
	usuarioManager.RegistrarUsuario(usuario)

	// Validation
	usuarioRegistrado := usuarioManager.GetUsuarioByUsername(username)

	if usuarioRegistrado == nil {
		t.Error("Expected usuario can't be nil")
	}
}

func TestUsuarioSinNombreNoSeRegistra(t *testing.T) {
	// Initialization
	usuarioManager := service.NewUsuarioManager()
	var usuario *domain.Usuario

	var nombre string = ""
	var username string = "francoagarcia"
	var mail string = "francoagarcia@gmail.com"
	var contrasenia string = "pass123"

	usuario = domain.NewUsuario(username, nombre, mail, contrasenia)

	// Operation
	err := usuarioManager.RegistrarUsuario(usuario)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "nombre is required" {
		t.Error("Expected error is nombre is required")
	}
}

func TestUsuarioSinUsernameNoSeRegistra(t *testing.T) {
	// Initialization
	usuarioManager := service.NewUsuarioManager()
	var usuario *domain.Usuario

	var nombre string = "Franco"
	var username string = ""
	var mail string = "francoagarcia@gmail.com"
	var contrasenia string = "pass123"

	usuario = domain.NewUsuario(username, nombre, mail, contrasenia)

	// Operation
	err := usuarioManager.RegistrarUsuario(usuario)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "username is required" {
		t.Error("Expected error is username is required")
	}
}

func TestUsuarioSinMailNoSeRegistra(t *testing.T) {
	// Initialization
	usuarioManager := service.NewUsuarioManager()
	var usuario *domain.Usuario

	var nombre string = "Franco"
	var username string = "francoagarcia"
	var mail string = ""
	var contrasenia string = "pass123"

	usuario = domain.NewUsuario(username, nombre, mail, contrasenia)

	// Operation
	err := usuarioManager.RegistrarUsuario(usuario)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "mail is required" {
		t.Error("Expected error is email is required")
	}
}

func TestUsuarioSinPasswordNoSeRegistra(t *testing.T) {
	// Initialization
	usuarioManager := service.NewUsuarioManager()
	var usuario *domain.Usuario

	var nombre string = "Franco"
	var username string = "francoagarcia"
	var mail string = "francoagarcia@gmail.com"
	var contrasenia string = ""

	usuario = domain.NewUsuario(username, nombre, mail, contrasenia)

	// Operation
	err := usuarioManager.RegistrarUsuario(usuario)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "contrasenia is required" {
		t.Error("Expected error is password is required")
	}
}
