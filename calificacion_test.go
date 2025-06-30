package tests

import (
	"streaming/models"
	"testing"
)

func TestNuevaCalificacionValida(t *testing.T) {
	usuario := models.NuevoUsuario("Ana", "ana@example.com", 28)
	contenido := models.NuevoContenido("Inception", "Ciencia Ficción", 2010, "Película")

	cal, err := models.NuevaCalificacion(usuario, contenido, 5, "Excelente película")

	if err != nil {
		t.Errorf("No se esperaba error para calificación válida: %v", err)
	}
	if cal == nil {
		t.Fatal("Se esperaba una calificación válida, se obtuvo nil")
	}
	if cal.Estrellas() != 5 {
		t.Errorf("Se esperaban 5 estrellas, se obtuvo %d", cal.Estrellas())
	}
}

func TestNuevaCalificacionInvalida(t *testing.T) {
	usuario := models.NuevoUsuario("Luis", "luis@example.com", 35)
	contenido := models.NuevoContenido("Avatar", "Fantasía", 2009, "Película")

	cal, err := models.NuevaCalificacion(usuario, contenido, 6, "Muy buena")

	if err == nil {
		t.Error("Se esperaba un error por estrellas fuera de rango, pero no ocurrió")
	}
	if cal != nil {
		t.Error("No se esperaba una calificación válida con estrellas > 5")
	}
}
