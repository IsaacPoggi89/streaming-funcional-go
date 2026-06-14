package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// ============================================================================
// CONTRACTS & INTERFACES (POLIMORFISMO DE LA UNIDAD 3)
// ============================================================================

type Autenticador interface {
	Registrar(u Usuario) error
}

type Catalogo interface {
	AgregarPelicula(p Pelicula)
	OrdenarPorNovedad() []Pelicula
}

type FiltroHistorial interface {
	AgregarRegistro(h HistorialVideo)
	SepararHistorial(email string) ([]HistorialVideo, []HistorialVideo)
}

// ============================================================================
// ESTRUCTURAS DE DATOS (MÓDULOS DE LA APLICACIÓN ENCAPSULADOS)
// ============================================================================

// --- MÓDULO 1: CONTROL DE ACCESO ---
type Usuario struct {
	email      string
	contrasena string
}

func (u *Usuario) SetEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("error: estructura de correo electrónico inválida")
	}
	u.email = email
	return nil
}

func (u *Usuario) GetEmail() string { return u.email }

func (u *Usuario) SetContrasena(pass string) error {
	if len(pass) < 6 {
		return errors.New("error: la contraseña debe tener al menos 6 caracteres por seguridad")
	}
	u.contrasena = pass
	return nil
}

type ServicioAcceso struct {
	usuarios []Usuario
}

func (s *ServicioAcceso) Registrar(u Usuario) error {
	for _, existente := range s.usuarios {
		if existente.GetEmail() == u.GetEmail() {
			return errors.New("error: el correo electrónico ya se encuentra registrado")
		}
	}
	s.usuarios = append(s.usuarios, u)
	return nil
}

// --- MÓDULO 2: CATÁLOGO DE PELÍCULAS ---
type Pelicula struct {
	titulo      string
	genero      string
	fechaSubida time.Time
}

func NuevaPelicula(titulo, genero string, fecha time.Time) (Pelicula, error) {
	if titulo == "" {
		return Pelicula{}, errors.New("error: el título de la película no puede estar vacío")
	}
	if genero == "" {
		return Pelicula{}, errors.New("error: el género de la película no puede estar vacío")
	}
	return Pelicula{titulo: titulo, genero: genero, fechaSubida: fecha}, nil
}

func (p *Pelicula) GetTitulo() string         { return p.titulo }
func (p *Pelicula) GetFechaSubida() time.Time { return p.fechaSubida }

type ServicioCatalogo struct {
	peliculas []Pelicula
}

func (s *ServicioCatalogo) AgregarPelicula(p Pelicula) {
	s.peliculas = append(s.peliculas, p)
}

func (s *ServicioCatalogo) OrdenarPorNovedad() []Pelicula {
	copia := make([]Pelicula, len(s.peliculas))
	copy(copia, s.peliculas)

	sort.Slice(copia, func(i, j int) bool {
		return copia[i].GetFechaSubida().After(copia[j].GetFechaSubida())
	})
	return copia
}

// --- MÓDULO 3: HISTORIAL DE VIDEO ---
type HistorialVideo struct {
	email        string
	idPelicula   string
	tiempoActual int
	tiempoTotal  int
}

func NuevoHistorial(email, id string, actual, total int) (HistorialVideo, error) {
	if total <= 0 {
		return HistorialVideo{}, errors.New("error: el tiempo total del video debe ser mayor a cero")
	}
	if actual < 0 || actual > total {
		return HistorialVideo{}, errors.New("error: el tiempo actual no puede ser negativo ni mayor al tiempo total")
	}
	if email == "" || id == "" {
		return HistorialVideo{}, errors.New("error: el email y el ID de la película son obligatorios")
	}
	return HistorialVideo{email: email, idPelicula: id, tiempoActual: actual, tiempoTotal: total}, nil
}

func (h *HistorialVideo) GetEmail() string       { return h.email }
func (h *HistorialVideo) GetTiempos() (int, int) { return h.tiempoActual, h.tiempoTotal }

type ServicioHistorial struct {
	registros []HistorialVideo
}

func (s *ServicioHistorial) AgregarRegistro(h HistorialVideo) {
	s.registros = append(s.registros, h)
}

func (s *ServicioHistorial) SepararHistorial(email string) ([]HistorialVideo, []HistorialVideo) {
	var seguirViendo []HistorialVideo
	var volverAVer []HistorialVideo

	for _, reg := range s.registros {
		if reg.GetEmail() == email {
			actual, total := reg.GetTiempos()
			if actual == total {
				volverAVer = append(volverAVer, reg)
			} else if actual < total {
				seguirViendo = append(seguirViendo, reg)
			}
		}
	}
	return seguirViendo, volverAVer
}

// --- ESTRUCTURAS DEL AUTÓNOMO 1 ---
type Canal struct {
	nombre string
	genero string
	activo bool
}

// ============================================================================
// FLUJO PRINCIPAL DE INTEGRACIÓN Y SIMULACIÓN
// ============================================================================
func main() {
	fmt.Println("=== SISTEMA DE STREAMING INTEGRADO ===")
	fmt.Println("Integración: Autónomo 1 (Estructuras) + Autónomo 2 (POO)")
	fmt.Println("-------------------------------------------------------")

	fmt.Println(">> [Autónomo 1] Configuración de Canales:")
	c1 := Canal{nombre: "Acción Total", genero: "Películas", activo: true}
	c2 := Canal{nombre: "Cine Documental", genero: "Cultura", activo: false}

	estado1 := "Inactivo"
	if c1.activo {
		estado1 = "Activo"
	}
	estado2 := "Inactivo"
	if c2.activo {
		estado2 = "Activo"
	}

	fmt.Printf("Canal: %s | Género: %s | Estado: %s\n", c1.nombre, c1.genero, estado1)
	fmt.Printf("Canal: %s | Género: %s | Estado: %s\n", c2.nombre, c2.genero, estado2)
	fmt.Println("-------------------------------------------------------")

	fmt.Println(">> [Autónomo 2] Lógica Funcional y de Historial:")

	var auth Autenticador = &ServicioAcceso{}
	var cine Catalogo = &ServicioCatalogo{}
	var tracker FiltroHistorial = &ServicioHistorial{}

	u := Usuario{}
	u.SetEmail("isaac@uide.edu.ec")
	u.SetContrasena("cyberseguridad2026")
	auth.Registrar(u)

	fmt.Println("[UIDE Play] Reproduciendo ahora: Introducción a la Ciberseguridad...")
	fmt.Printf("¡'Introducción a la Ciberseguridad' se ha añadido al historial de %s!\n", "Isaac")
	fmt.Printf("¡'Programación Avanzada en Go' se ha añadido al historial de %s!\n", "Isaac")

	p1, _ := NuevaPelicula("Introducción a la Ciberseguridad", "Estudio", time.Now().Add(-24*time.Hour))
	p2, _ := NuevaPelicula("Programación Avanzada en Go", "Estudio", time.Now())
	cine.AgregarPelicula(p1)
	cine.AgregarPelicula(p2)

	h1, _ := NuevoHistorial("isaac@uide.edu.ec", "CiberID", 2700, 2700)
	h2, _ := NuevoHistorial("isaac@uide.edu.ec", "GoID", 3600, 3600)
	tracker.AgregarRegistro(h1)
	tracker.AgregarRegistro(h2)

	_, volverAVer := tracker.SepararHistorial("isaac@uide.edu.ec")

	fmt.Println("\n--- Historial de Reproducción de Isaac ---")
	contador := 1
	for _, vv := range volverAVer {
		for _, pelicula := range cine.OrdenarPorNovedad() {
			if (pelicula.GetTitulo() == "Introducción a la Ciberseguridad" && vv.idPelicula == "CiberID") ||
				(pelicula.GetTitulo() == "Programación Avanzada en Go" && vv.idPelicula == "GoID") {
				_, tTotal := vv.GetTiempos()
				fmt.Printf("%d. %s (%d min)\n", contador, pelicula.GetTitulo(), tTotal/60)
				contador++
				break
			}
		}
	}
}
