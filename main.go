package main

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("La suma de 5 + 5 es:", Sumar(5, 5))
}
