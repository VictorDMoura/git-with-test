package iteracao

func Repetir(caractere string, qtd int) string {
	var repeticoes string
	for i := 0; i < qtd; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
