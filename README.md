# Sistema de Gestión de Streaming en Go

## Descripción

Este proyecto es un sistema básico de gestión de streaming desarrollado en Go. Permite registrar usuarios, agregar contenidos (películas y series), reproducir contenido de forma simulada y agregar calificaciones. Además, ofrece servicios web REST que permiten interactuar con el sistema mediante JSON.

## Funcionalidades principales

- Registro y validación de usuarios (prevención de correos duplicados).
- Gestión de contenidos con metadatos: título, género, año y tipo (película o serie).
- Simulación de reproducción de contenidos.
- Calificación de contenidos con validación de estrellas (1 a 5) y comentarios.
- Listado de usuarios, contenidos y calificaciones.
- Servicios web REST para operaciones CRUD sobre usuarios, contenidos y calificaciones.
- Manejo de concurrencia para atender múltiples solicitudes HTTP.

## Tecnologías usadas

- Lenguaje: Go (Golang)
- Framework web: net/http y gorilla/mux
- Control de versiones: Git y GitHub

## Estructura del repositorio
