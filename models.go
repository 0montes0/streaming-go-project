package main

import (
    "errors"
    "strings"
)

type Usuario struct {
    Nombre string `json:"nombre"`
    Correo string `json:"correo"`
    Edad   int    `json:"edad"`
}

type Contenido struct {
    Titulo string `json:"titulo"`
    Genero string `json:"genero"`
    Anio   int    `json:"anio"`
    Tipo   string `json:"tipo"`
}

type Calificacion struct {
    Usuario    *Usuario `json:"usuario,omitempty"`
    Contenido  *Contenido `json:"contenido,omitempty"`
    Estrellas  int      `json:"estrellas"`
    Comentario string   `json:"comentario"`
}

type SistemaStreaming struct {
    Usuarios       []*Usuario
    Contenidos     []*Contenido
    Calificaciones []*Calificacion
}

func NuevoSistema() *SistemaStreaming {
    return &SistemaStreaming{}
}

func (s *SistemaStreaming) AgregarUsuario(nombre, correo string, edad int) error {
    correo = strings.ToLower(correo)
    for _, u := range s.Usuarios {
        if u.Correo == correo {
            return errors.New("el correo ya está registrado")
        }
    }
    s.Usuarios = append(s.Usuarios, &Usuario{Nombre: nombre, Correo: correo, Edad: edad})
    return nil
}

func (s *SistemaStreaming) AgregarContenido(titulo, genero string, anio int, tipo string) {
    s.Contenidos = append(s.Contenidos, &Contenido{Titulo: titulo, Genero: genero, Anio: anio, Tipo: tipo})
}

func (s *SistemaStreaming) BuscarContenido(titulo string) *Contenido {
    for _, c := range s.Contenidos {
        if c.Titulo == titulo {
            return c
        }
    }
    return nil
}

func (s *SistemaStreaming) BuscarUsuario(correo string) (*Usuario, error) {
    correo = strings.ToLower(correo)
    for _, u := range s.Usuarios {
        if u.Correo == correo {
            return u, nil
        }
    }
    return nil, errors.New("usuario no encontrado")
}

func (s *SistemaStreaming) AgregarCalificacion(correoUsuario, tituloContenido string, estrellas int, comentario string) error {
    if estrellas < 1 || estrellas > 5 {
        return errors.New("las estrellas deben estar entre 1 y 5")
    }

    u, err := s.BuscarUsuario(correoUsuario)
    if err != nil {
        return err
    }
    c := s.BuscarContenido(tituloContenido)
    if c == nil {
        return errors.New("contenido no encontrado")
    }

    cal := &Calificacion{
        Usuario:    u,
        Contenido:  c,
        Estrellas:  estrellas,
        Comentario: comentario,
    }
    s.Calificaciones = append(s.Calificaciones, cal)
    return nil
}
