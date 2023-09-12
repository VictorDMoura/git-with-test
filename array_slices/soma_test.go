package arrayslices

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		verificarResultado(t, resultado, esperado, numeros)
	})

}

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v esperado %v", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {
	resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
	esperado := []int{2, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v esperado %v", resultado, esperado)
	}
}

func verificarResultado(t testing.TB, resultado, esperado int, numeros []int) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%d', esperado '%d', dado %v", resultado, esperado, numeros)
	}
}
