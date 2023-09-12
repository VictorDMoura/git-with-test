package arrayslices

import "testing"

func TestSoma(t *testing.T) {

	numeros := [5]int{1, 2, 3, 4, 5}

	resultado := Soma(numeros)
	espereado := 15

	if resultado != espereado {
		t.Errorf("resultado '%d', esperado '%d', dado %v", resultado, espereado, numeros)
	}
}
