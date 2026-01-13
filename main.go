package main

import (
	"fmt"
	"net/http"
)

func Sumar(a, b int) int {
	return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hola Mundo</h1>\n")
	fmt.Fprintf(w, "<p>La suma de 5 + 5 es: <strong>%d</strong></p>\n", Sumar(5, 5))
}

func iniciarServidor() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8089...")
	http.ListenAndServe(":8089", nil)
}

func main() {
	iniciarServidor()
}
