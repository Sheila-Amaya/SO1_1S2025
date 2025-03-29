package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo, usando NodePort en Kubernetes – Primer Semestre 2025\nCarnet: 202000558")
}

func main() {
	port := "4000"
	http.HandleFunc("/", handler)
	fmt.Println("Servidor corriendo en http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
