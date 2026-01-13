package main

import (
	"fmt"
	"net/http"
)

func Sumar(a, b int) int {
	return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ruber in Go v2\n")
	fmt.Fprintf(w, "La suma de 5 + 5 es: %d\n", Sumar(5, 5))
}

func iniciarServidor() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func main() {
	iniciarServidor()
}
