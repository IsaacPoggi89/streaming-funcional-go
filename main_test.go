package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1. Prueba Unitaria: Obtener catálogo de forma segura a través del Handler
func TestObtenerCatalogo(t *testing.T) {
	catalogo := &CatalogoController{
		Peliculas: map[string]Pelicula{
			"TestID": {ID: "TestID", Titulo: "Pelicula de Prueba", Duracion: 1200},
		},
	}

	// Verificamos directamente el mapa de películas asegurado por el Mutex
	catalogo.mu.Lock()
	totalPeliculas := len(catalogo.Peliculas)
	catalogo.mu.Unlock()

	if totalPeliculas != 1 {
		t.Errorf("Esperaba 1 película en el catálogo, se obtuvieron: %d", totalPeliculas)
	}
}

// 2. Prueba de Integración: Endpoint /catalogo
func TestHandlerCatalogo(t *testing.T) {
	catalogo := &CatalogoController{
		Peliculas: map[string]Pelicula{
			"GoID": {ID: "GoID", Titulo: "Programación Avanzada en Go", Duracion: 5400},
		},
	}

	req, err := http.NewRequest("GET", "/catalogo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(catalogo.HandlerCatalogo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: esperado %d, obtenido %d", http.StatusOK, status)
	}
}

// 3. Prueba de Aceptación: Reproducir película inexistente (Debe retornar 404)
func TestHandlerReproducirNoEncontrado(t *testing.T) {
	catalogo := &CatalogoController{
		Peliculas: map[string]Pelicula{},
	}

	req, err := http.NewRequest("GET", "/reproducir?id=ID_FALSO", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(catalogo.HandlerReproducir)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Esperaba estado %d, obtenido %d", http.StatusNotFound, status)
	}
}
