package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Estructura que representa una Película (Modelo/Objeto)
type Pelicula struct {
	ID       string `json:"id"`
	Titulo   string `json:"titulo"`
	Duracion int    `json:"duracion_segundos"`
}

// Estructura para manejar el catálogo y aplicar concurrencia segura
type CatalogoController struct {
	mu        sync.Mutex
	Peliculas map[string]Pelicula
}

// Método orientado a objetos para obtener el catálogo
func (c *CatalogoController) ObtenerCatalogo() []Pelicula {
	c.mu.Lock()
	defer c.mu.Unlock()

	lista := []Pelicula{}
	for _, p := range c.Peliculas {
		lista = append(lista, p)
	}
	return lista
}

// Controlador para la ruta del catálogo usando concurrencia nativa
func (c *CatalogoController) HandlerCatalogo(w http.ResponseWriter, r *http.Request) {
	go log.Printf("Petición concurrente recibida desde: %s", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	catalogo := c.ObtenerCatalogo()
	json.NewEncoder(w).Encode(catalogo)
}

// Simulación de reproducción que demuestra concurrencia y POO
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

	// Simulación de transmisión concurrente desacoplada (Goroutine)
	go func(p Pelicula) {
		fmt.Printf("[Streaming POO] Iniciando reproducción asíncrona de: %s\n", p.Titulo)
		time.Sleep(2 * time.Second)
		fmt.Printf("[Streaming POO] Emisión completada de: %s\n", p.Titulo)
	}(peli)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje":      "Reproducción iniciada exitosamente",
		"titulo":       peli.Titulo,
		"concurrencia": "Ejecutándose en segundo plano (Goroutine)",
	})
}

func main() {
	// Inicialización del catálogo
	catalogo := &CatalogoController{
		Peliculas: map[string]Pelicula{
			"CiberID": {ID: "CiberID", Titulo: "Introducción a la Ciberseguridad", Duracion: 3600},
			"GoID":    {ID: "GoID", Titulo: "Programación Avanzada en Go", Duracion: 5400},
		},
	}

	http.HandleFunc("/catalogo", catalogo.HandlerCatalogo)
	http.HandleFunc("/reproducir", catalogo.HandlerReproducir)

	fmt.Println("Servidor Web de Streaming escuchando en http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
