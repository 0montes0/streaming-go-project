// streaming.go
// Proyecto: Sistema de Gestión de Streaming en Go
// Autor: Andrés Montesdeoca
// Fecha: Junio 2025
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type usuario struct {
	nombre string
	correo string
	edad   int
}

func nuevoUsuario(nombre, correo string, edad int) *usuario {
	return &usuario{nombre: nombre, correo: strings.ToLower(correo), edad: edad}
}

func (u *usuario) Nombre() string { return u.nombre }
func (u *usuario) Correo() string { return u.correo }
func (u *usuario) Edad() int      { return u.edad }

type contenido struct {
	titulo string
	genero string
	anio   int
	tipo   string
}

func nuevoContenido(titulo, genero string, anio int, tipo string) *contenido {
	return &contenido{titulo: titulo, genero: genero, anio: anio, tipo: tipo}
}

func (c *contenido) MostrarInfo() {
	fmt.Printf("Título: %s | Género: %s | Año: %d | Tipo: %s\n", c.titulo, c.genero, c.anio, c.tipo)
}

func (c *contenido) Reproducir() {
	fmt.Printf("Reproduciendo %s: %s\n", c.tipo, c.titulo)
}

type calificacion struct {
	usuario    *usuario
	contenido  *contenido
	estrellas  int
	comentario string
}

func nuevaCalificacion(u *usuario, c *contenido, estrellas int, comentario string) (*calificacion, error) {
	if estrellas < 1 || estrellas > 5 {
		return nil, errors.New("las estrellas deben estar entre 1 y 5")
	}
	return &calificacion{usuario: u, contenido: c, estrellas: estrellas, comentario: comentario}, nil
}

func (cal *calificacion) Mostrar() {
	fmt.Printf("Usuario: %s calificó '%s' con %d estrellas. Comentario: %s\n",
		cal.usuario.Nombre(), cal.contenido.titulo, cal.estrellas, cal.comentario)
}

type sistemaStreaming struct {
	usuarios       []*usuario
	contenidos     []*contenido
	calificaciones []*calificacion
}

func nuevoSistema() *sistemaStreaming {
	return &sistemaStreaming{}
}

func (s *sistemaStreaming) AgregarUsuario(nombre, correo string, edad int) error {
	for _, u := range s.usuarios {
		if u.Correo() == strings.ToLower(correo) {
			return errors.New("el correo ya está registrado")
		}
	}
	s.usuarios = append(s.usuarios, nuevoUsuario(nombre, correo, edad))
	return nil
}

func (s *sistemaStreaming) ListarUsuarios() {
	fmt.Println("Usuarios registrados:")
	for _, u := range s.usuarios {
		fmt.Printf("- %s (%s), Edad: %d\n", u.Nombre(), u.Correo(), u.Edad())
	}
}

func (s *sistemaStreaming) AgregarContenido(titulo, genero string, anio int, tipo string) {
	s.contenidos = append(s.contenidos, nuevoContenido(titulo, genero, anio, tipo))
}

func (s *sistemaStreaming) ListarContenidos() {
	fmt.Println("Catálogo de contenido:")
	for _, c := range s.contenidos {
		c.MostrarInfo()
	}
}

func (s *sistemaStreaming) BuscarUsuario(correo string) (*usuario, error) {
	for _, u := range s.usuarios {
		if u.Correo() == strings.ToLower(correo) {
			return u, nil
		}
	}
	return nil, errors.New("usuario no encontrado")
}

func (s *sistemaStreaming) ReproducirContenido() {
	fmt.Println("Reproduciendo contenido:")
	for _, c := range s.contenidos {
		c.Reproducir()
	}
}

func (s *sistemaStreaming) AgregarCalificacion(correoUsuario, tituloContenido string, estrellas int, comentario string) error {
	u, err := s.BuscarUsuario(correoUsuario)
	if err != nil {
		return err
	}

	var cont *contenido
	for _, c := range s.contenidos {
		if c.titulo == tituloContenido {
			cont = c
			break
		}
	}
	if cont == nil {
		return errors.New("contenido no encontrado")
	}

	cal, err := nuevaCalificacion(u, cont, estrellas, comentario)
	if err != nil {
		return err
	}
	s.calificaciones = append(s.calificaciones, cal)
	return nil
}

func (s *sistemaStreaming) MostrarCalificaciones() {
	fmt.Println("Calificaciones:")
	for _, cal := range s.calificaciones {
		cal.Mostrar()
	}
}

func leerTexto(msg string) string {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerEntero(msg string) int {
	texto := leerTexto(msg)
	num, _ := strconv.Atoi(texto)
	return num
}

func mostrarMenu() {
	fmt.Println("\n--- Menú del Sistema de Streaming ---")
	fmt.Println("1. Agregar usuario")
	fmt.Println("2. Agregar contenido")
	fmt.Println("3. Ver usuarios")
	fmt.Println("4. Ver contenidos")
	fmt.Println("5. Reproducir contenido")
	fmt.Println("6. Calificar contenido")
	fmt.Println("7. Ver calificaciones")
	fmt.Println("0. Salir")
}

func main() {
	sistema := nuevoSistema()
	for {
		mostrarMenu()
		opcion := leerEntero("Elige una opción: ")

		switch opcion {
		case 1:
			nombre := leerTexto("Nombre: ")
			correo := leerTexto("Correo: ")
			edad := leerEntero("Edad: ")
			if err := sistema.AgregarUsuario(nombre, correo, edad); err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			titulo := leerTexto("Título: ")
			genero := leerTexto("Género: ")
			anio := leerEntero("Año: ")
			tipo := leerTexto("Tipo (Película/Serie): ")
			sistema.AgregarContenido(titulo, genero, anio, tipo)
		case 3:
			sistema.ListarUsuarios()
		case 4:
			sistema.ListarContenidos()
		case 5:
			sistema.ReproducirContenido()
		case 6:
			correo := leerTexto("Correo del usuario: ")
			titulo := leerTexto("Título del contenido: ")
			estrellas := leerEntero("Estrellas (1-5): ")
			comentario := leerTexto("Comentario: ")
			if err := sistema.AgregarCalificacion(correo, titulo, estrellas, comentario); err != nil {
				fmt.Println("Error:", err)
			}
		case 7:
			sistema.MostrarCalificaciones()
		case 0:
			fmt.Println("Gracias por usar el sistema.")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
