package main

import "testing"

func TestSumar(t *testing.T) {
	resultado := Sumar(2, 3)
	esperado := 5

	if resultado != esperado {
		t.Errorf("Resultado incorrecto: obtuvo %d, esperaba %d", resultado, esperado)
	}
}
