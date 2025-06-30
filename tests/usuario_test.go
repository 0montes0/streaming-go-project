// tests/usuario_test.go
package tests

import (
	"streaming/models"
	"testing"
)

func TestNuevoUsuario(t *testing.T) {
	u := models.NuevoUsuario("Juan Perez", "juan@example.com", 30)
	if u.Nombre() != "Juan Perez" {
		t.Errorf("Esperado 'Juan Perez', pero se obtuvo '%s'", u.Nombre())
	}
	if u.Correo() != "juan@example.com" {
		t.Errorf("Esperado 'juan@example.com', pero se obtuvo '%s'", u.Correo())
	}
	if u.Edad() != 30 {
		t.Errorf("Esperado 30, pero se obtuvo %d", u.Edad())
	}
}
