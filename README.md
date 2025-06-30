# Sistema de Gestión de Streaming en Go

## Descripción
Este proyecto es un sistema completo de gestión de streaming desarrollado en Go. Permite registrar usuarios, administrar contenidos (películas y series), reproducir contenido y agregar calificaciones. Incluye una interfaz de consola y servicios web REST con JSON.

## Funcionalidades principales
- Registro y gestión de usuarios con validación.
- Gestión de catálogo de contenidos.
- Reproducción simulada de videos.
- Calificación de contenidos con estrellas y comentarios.
- Servicios web REST para integración y acceso remoto.
- Manejo de concurrencia para atender múltiples solicitudes.

## Estructura del proyecto
- `streaming.go`: lógica y menú consola (no incluido en esta versión modular).
- `main.go`: servidor HTTP con servicios web REST.
- `handlers.go`: manejadores HTTP para las rutas.
- `models.go`: definiciones de structs y métodos OOP.
- `tests/`: pruebas unitarias.
- `docs/`: documentación y diagramas.
- `videos/`: video demostrativo.

## Tecnologías usadas
- Lenguaje: Go (Golang)
- Paquetes: Gorilla Mux para routing HTTP.
- Control de versiones: GitHub.

## Cómo ejecutar el proyecto
1. Clona el repositorio:
