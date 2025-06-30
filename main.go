package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Rutas para servicios web
    r.HandleFunc("/usuarios", GetUsuarios).Methods("GET")
    r.HandleFunc("/usuarios", CrearUsuario).Methods("POST")
    r.HandleFunc("/contenidos", GetContenidos).Methods("GET")
    r.HandleFunc("/contenidos", CrearContenido).Methods("POST")
    r.HandleFunc("/stream", StreamVideo).Methods("GET")
    r.HandleFunc("/calificaciones", GetCalificaciones).Methods("GET")
    r.HandleFunc("/calificaciones", CrearCalificacion).Methods("POST")

    log.Println("Servidor corriendo en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
