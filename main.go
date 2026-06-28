package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Estructura: Película (Modelo/Objeto)
type Pelicula struct {
	ID       string `json:"id"`
	Titulo   string `json:"titulo"`
	Duracion int    `json:"duracion_segundos"`
}

// Estructura: Usuario (Modelo/Objeto)
type Usuario struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Membrezia string `json:"tipo_membrezia"`
}

// Controlador para manejar el catálogo de forma concurrente y segura
type CatalogoController struct {
	mu        sync.Mutex
	Peliculas map[string]Pelicula
}

// 1. Servicio Web: Consultar catálogo completo
func (c *CatalogoController) HandlerCatalogo(w http.ResponseWriter, r *http.Request) {
	go log.Printf("Petición concurrente recibida en /catalogo desde: %s", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")

	c.mu.Lock()
	lista := []Pelicula{}
	for _, p := range c.Peliculas {
		lista = append(lista, p)
	}
	c.mu.Unlock()

	json.NewEncoder(w).Encode(lista)
}

// 2. Servicio Web: Reproducir video asíncrono (Concurrencia)
func (c *CatalogoController) HandlerReproducir(w http.ResponseWriter, r *http.Request) {
	idPeli := r.URL.Query().Get("id")

	c.mu.Lock()
	peli, existe := c.Peliculas[idPeli]
	c.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	if !existe {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Película no encontrada"})
		return
	}

	// Simulación de emisión en segundo plano (Goroutine)
	go func(p Pelicula) {
		fmt.Printf("[Streaming POO] Iniciando transmisión asíncrona de: %s\n", p.Titulo)
		time.Sleep(2 * time.Second)
		fmt.Printf("[Streaming POO] Emisión completada de: %s\n", p.Titulo)
	}(peli)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje":      "Reproducción iniciada exitosamente",
		"titulo":       peli.Titulo,
		"concurrencia": "Ejecutándose en segundo plano (Goroutine)",
	})
}

// 3. Servicio Web: Simulación para agregar película
func (c *CatalogoController) HandlerAgregar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.mu.Lock()
	c.Peliculas["NewID"] = Pelicula{ID: "NewID", Titulo: "Contenido Agregado Dinámicamente", Duracion: 1800}
	c.mu.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"status": "Película agregada al catálogo"})
}

// 4. Servicio Web: Simulación para eliminar película
func (c *CatalogoController) HandlerEliminar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.mu.Lock()
	delete(c.Peliculas, "CiberID")
	c.mu.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"status": "Contenido eliminado del catálogo"})
}

// 5. Servicio Web: Buscar película
func (c *CatalogoController) HandlerBuscar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idPeli := r.URL.Query().Get("id")

	c.mu.Lock()
	peli, existe := c.Peliculas[idPeli]
	c.mu.Unlock()

	if !existe {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Contenido no hallado"})
		return
	}
	json.NewEncoder(w).Encode(peli)
}

// 6. Servicio Web: Gestión de perfil de usuario
func HandlerUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuario := Usuario{ID: "U123", Nombre: "Isaac Poggi", Membrezia: "Premium"}
	json.NewEncoder(w).Encode(usuario)
}

// 7. Servicio Web: Consultar historial de reproducción
func HandlerHistorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	historial := []string{"Programación Avanzada en Go", "Introducción a la Ciberseguridad"}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"usuario":   "Isaac Poggi",
		"historial": historial,
	})
}

// 8. Servicio Web: Monitoreo de actividad del servidor
func HandlerEstado(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"servidor":     "Streaming-Backend-POO",
		"estado":       "Online y Operativo",
		"concurrencia": "Habilitada (Goroutines activas)",
		"tiempo_up":    time.Now().Format(time.RFC3339),
	})
}

func main() {
	// Inicialización del catálogo con POO
	catalogo := &CatalogoController{
		Peliculas: map[string]Pelicula{
			"CiberID": {ID: "CiberID", Titulo: "Introducción a la Ciberseguridad", Duracion: 3600},
			"GoID":    {ID: "GoID", Titulo: "Programación Avanzada en Go", Duracion: 5400},
		},
	}

	// Mapeo de los 8 Servicios Web (End-points) solicitados en la rúbrica
	http.HandleFunc("/catalogo", catalogo.HandlerCatalogo)
	http.HandleFunc("/reproducir", catalogo.HandlerReproducir)
	http.HandleFunc("/agregar-pelicula", catalogo.HandlerAgregar)
	http.HandleFunc("/eliminar-pelicula", catalogo.HandlerEliminar)
	http.HandleFunc("/buscar-pelicula", catalogo.HandlerBuscar)
	http.HandleFunc("/perfil-usuario", HandlerUsuario)
	http.HandleFunc("/historial", HandlerHistorial)
	http.HandleFunc("/estado-servidor", HandlerEstado)

	fmt.Println("Servidor Web de Streaming escuchando en http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
