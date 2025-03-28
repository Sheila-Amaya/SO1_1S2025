package main

import (
	"fmt"
	"net/http"
)

// Ruta raíz "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "RAIZ")
}

// Ruta "/gestion"
func gestionHandler(w http.ResponseWriter, r *http.Request) {
	message := `
	<html>
		<head><title>Gestión</title></head>
		<body>
			<h1>Sistemas Operativos 1 - Primer Semestre 2025</h1>
			<h2>Carnet: 202000558</h2>
		</body>
	</html>
	`
	fmt.Fprintln(w, message)
}

func main() {
	port := "4000"

	// Definir las rutas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/gestion", gestionHandler)

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
