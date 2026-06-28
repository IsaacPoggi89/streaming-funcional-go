# streaming-funcional-go
Etapa 1: Planeación y lógica funcional para sistema de streaming
##  Descripción del Proyecto
Este repositorio contiene la planeación y el diseño lógico de una plataforma de streaming de video automatizada. El sistema está proyectado para gestionar el acceso de usuarios, administrar un catálogo de contenidos audiovisuales y personalizar la experiencia del espectador mediante la segmentación inteligente de su historial de reproducción.
##  Paradigma de Programación
El diseño lógico de este sistema se fundamenta en los principios de la **programación funcional**. Se prioriza el uso de funciones puras, colecciones de datos inmutables y operaciones que evitan efectos secundarios, garantizando un entorno predecible, escalable y seguro para el procesamiento de datos concurrentes.

##  Estructura de Módulos Planeados
Para cumplir con los objetivos del sistema, la arquitectura se divide en tres componentes conceptuales:
* **Control de Acceso:** Gestión púra de autenticación, registro inmutable de cuentas y validación de credenciales.
* **Catálogo de Películas:** Administración y ordenamiento cronológico de los contenidos audiovisuales disponibles.
* **Historial de Reproducción:** Segmentación matemática del progreso del usuario para dividir contenidos entre "Continuar viendo" y "Ver de nuevo".

## 👥 Datos del Proyecto
* **Modalidad:** Trabajo Individual
* **Autor:** Isaac Poggi


# Plataforma de Streaming de Video - Proyecto Final POO

## Datos del Grupo
* **Estudiante(s):** Isaac Poggi
* **Materia:** Programación Orientada a Objetos
* **Fecha:** 27 de Junio de 2026

## Objetivo del Programa
Evolucionar e integrar los conocimientos de las 8 semanas (4 unidades) de la materia, transformando un sistema de consola básico en una arquitectura de backend web concurrente orientada a objetos. El sistema permite administrar catálogos y gestionar la reproducción multimedia de forma asíncrona y segura ante múltiples peticiones.

## Principales Funcionalidades del Código
1. **Modelo Orientado a Objetos:** Encapsulamiento de catálogos, usuarios y contenidos multimedia mediante estructuras (`structs`) y métodos asociados en Go.
2. **Servidor Web Concurrente:** Implementación de la librería nativa `net/http` para despachar peticiones de manera asíncrona mediante *goroutines* (hilos ligeros), evitando el bloqueo del servidor.
3. **Gestión Segura de Concurrencia:** Uso de bloqueos con `sync.Mutex` para evitar condiciones de carrera (*race conditions*) al acceder a la memoria del catálogo de películas.
4. **Serialización JSON:** Comunicación de datos cliente-servidor estructurada mediante formato estándar JSON.
5. **Pruebas Automatizadas:** Validación de robustez, integración y aceptación implementadas con el paquete nativo `testing`.
