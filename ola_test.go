package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Victor")
	esperado := "Olá, Victor!"

	if resultado != esperado {
		t.Errorf("resultado '%q', esperado '%q'", resultado, esperado)
	}
}
