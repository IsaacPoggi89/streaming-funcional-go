package main

import "fmt"

// ==========================================
// 1. ESTRUCTURAS DEL AUTÓNOMO 1 (Canales de Streaming)
// ==========================================

// Canal representa la estructura base que definiste en el primer trabajo
type Canal struct {
	Nombre string
	Genero string
	Activo bool
}

// MostrarInformacionCanal es la función del Autónomo 1 para desplegar los datos del canal
func MostrarInformacionCanal(c Canal) {
	estado := "Inactivo"
	if c.Activo {
		estado = "Activo"
	}
	fmt.Printf("Canal: %s | Género: %s | Estado: %s\n", c.Nombre, c.Genero, estado)
}

// ==========================================
// 2. MÓDULO DE HISTORIAL & ENCAPSULACIÓN (Autónomo 2)
// ==========================================

// Video representa la estructura de un contenido audiovisual.
// Las propiedades están en minúscula (privadas) para cumplir con la Encapsulación.
type Video struct {
	titulo   string
	duracion int // en minutos
}

// NewVideo es el constructor para crear un nuevo video de forma segura.
func NewVideo(titulo string, duracion int) Video {
	return Video{titulo: titulo, duracion: duracion}
}

// GetTitulo permite acceder al título del video desde fuera de la estructura.
func (v Video) GetTitulo() string {
	return v.titulo
}

// HistorialUsuario gestiona la lista de videos reproducidos por un usuario.
type HistorialUsuario struct {
	Usuario string
	Videos  []Video
}

// AgregarVideo añade un video visitado al historial del usuario.
func (h *HistorialUsuario) AgregarVideo(v Video) {
	h.Videos = append(h.Videos, v)
	fmt.Printf("¡'%s' se ha añadido al historial de %s!\n", v.GetTitulo(), h.Usuario)
}

// MostrarHistorial imprime en pantalla todas las reproducciones de forma ordenada.
func (h HistorialUsuario) MostrarHistorial() {
	fmt.Printf("\n--- Historial de Reproducción de %s ---\n", h.Usuario)
	if len(h.Videos) == 0 {
		fmt.Println("El historial está vacío.")
		return
	}
	for i, video := range h.Videos {
		fmt.Printf("%d. %s (%d min)\n", i+1, video.GetTitulo(), video.duracion)
	}
}

// ==========================================
// 3. INTERFACES (Autónomo 2)
// ==========================================

// Reproductor define el comportamiento de cualquier sistema que reproduzca contenido.
type Reproductor interface {
	Reproducir()
}

// ServicioStreaming implementa la interfaz Reproductor.
type ServicioStreaming struct {
	NombrePlataforma string
	VideoActual      Video
}

// Reproducir ejecuta la acción requerida por la interfaz Reproductor.
func (s ServicioStreaming) Reproducir() {
	fmt.Printf("[%s] Reproduciendo ahora: %s...\n", s.NombrePlataforma, s.VideoActual.GetTitulo())
}

// ==========================================
// 4. EJECUCIÓN PRINCIPAL (Integración Total)
// ==========================================

func main() {
	fmt.Println("=== SISTEMA DE STREAMING INTEGRADO ===")
	fmt.Println("Integración: Autónomo 1 (Estructuras) + Autónomo 2 (POO)")
	fmt.Println("-----------------------------------------------------")

	// --- PARTE 1: Datos del Autónomo 1 ---
	fmt.Println(">> [Autónomo 1] Configuración de Canales:")
	canal1 := Canal{Nombre: "Acción Total", Genero: "Películas", Activo: true}
	canal2 := Canal{Nombre: "Cine Documental", Genero: "Cultura", Activo: false}

	MostrarInformacionCanal(canal1)
	MostrarInformacionCanal(canal2)
	fmt.Println("-----------------------------------------------------")

	// --- PARTE 2: Datos del Autónomo 2 ---
	fmt.Println(">> [Autónomo 2] Lógica Funcional y de Historial:")

	// 1. Crear videos usando el constructor seguro (Encapsulación)
	video1 := NewVideo("Introducción a la Ciberseguridad", 45)
	video2 := NewVideo("Programación Avanzada en Go", 60)

	// 2. Inicializar el historial del usuario
	historialIsaac := HistorialUsuario{
		Usuario: "Isaac",
	}

	// 3. Usar la Interface para simular la reproducción en la plataforma
	plataforma := ServicioStreaming{
		NombrePlataforma: "UIDE Play",
		VideoActual:      video1,
	}

	// Ejecución mediante la abstracción de la interfaz
	var reproductor Reproductor = plataforma
	reproductor.Reproducir()

	// 4. Alimentar el historial del usuario conforme consume contenido
	historialIsaac.AgregarVideo(video1)
	historialIsaac.AgregarVideo(video2)

	// 5. Desplegar los resultados finales del avance de usuario
	historialIsaac.MostrarHistorial()
}
