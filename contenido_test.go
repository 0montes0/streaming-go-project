package tests

import (
	"streaming/models"
	"testing"
)

func TestNuevoContenido(t *testing.T) {
	titulo := "Matrix"
	genero := "Ciencia Ficción"
	anio := 1999
	tipo := "Película"

	contenido := models.NuevoContenido(titulo, genero, anio, tipo)

	if contenido == nil {
		t.Fatal("El contenido no debería ser nil")
	}

	if contenido.Titulo() != titulo {
		t.Errorf("Se esperaba título '%s', pero se obtuvo '%s'", titulo, contenido.Titulo())
	}
	if contenido.Genero() != genero {
		t.Errorf("Se esperaba género '%s', pero se obtuvo '%s'", genero, contenido.Genero())
	}
	if contenido.Anio() != anio {
		t.Errorf("Se esperaba año '%d', pero se obtuvo '%d'", anio, contenido.Anio())
	}
	if contenido.Tipo() != tipo {
		t.Errorf("Se esperaba tipo '%s', pero se obtuvo '%s'", tipo, contenido.Tipo())
	}
}
