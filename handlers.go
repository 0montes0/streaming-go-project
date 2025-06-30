package main

import (
    "encoding/json"
    "net/http"
    "strconv"
)

var sistema = NuevoSistema() // Instancia global para manejar datos

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sistema.Usuarios)
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
    var u Usuario
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        http.Error(w, "Entrada inválida", http.StatusBadRequest)
        return
    }
    err = sistema.AgregarUsuario(u.Nombre, u.Correo, u.Edad)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(u)
}

func GetContenidos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sistema.Contenidos)
}

func CrearContenido(w http.ResponseWriter, r *http.Request) {
    var c Contenido
    err := json.NewDecoder(r.Body).Decode(&c)
    if err != nil {
        http.Error(w, "Entrada inválida", http.StatusBadRequest)
        return
    }
    sistema.AgregarContenido(c.Titulo, c.Genero, c.Anio, c.Tipo)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(c)
}

func StreamVideo(w http.ResponseWriter, r *http.Request) {
    titulo := r.URL.Query().Get("titulo")
    if titulo == "" {
        http.Error(w, "Falta parámetro 'titulo'", http.StatusBadRequest)
        return
    }
    contenido := sistema.BuscarContenido(titulo)
    if contenido == nil {
        http.Error(w, "Contenido no encontrado", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Reproduciendo: " + contenido.Titulo))
}

func GetCalificaciones(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sistema.Calificaciones)
}

func CrearCalificacion(w http.ResponseWriter, r *http.Request) {
    var cal CalificacionInput
    err := json.NewDecoder(r.Body).Decode(&cal)
    if err != nil {
        http.Error(w, "Entrada inválida", http.StatusBadRequest)
        return
    }
    err = sistema.AgregarCalificacion(cal.CorreoUsuario, cal.TituloContenido, cal.Estrellas, cal.Comentario)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cal)
}

type CalificacionInput struct {
    CorreoUsuario  string `json:"correo_usuario"`
    TituloContenido string `json:"titulo_contenido"`
    Estrellas      int    `json:"estrellas"`
    Comentario     string `json:"comentario"`
}
