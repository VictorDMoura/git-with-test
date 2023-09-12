package iteracao

import (
	"strings"
	"testing"
)

const quantidadeDeLoops = 10
const palavraRepetir = "a"

func TestRepetir(t *testing.T) {
	for i := 0; i < quantidadeDeLoops; i++ {
		repeticoes := Repetir(palavraRepetir, i)
		esperado := strings.Repeat(palavraRepetir, i)
		if repeticoes != esperado {
			t.Errorf("Esperado '%q', resultado '%q'", esperado, repeticoes)
		}
	}

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir(palavraRepetir, quantidadeDeLoops)
	}
}
