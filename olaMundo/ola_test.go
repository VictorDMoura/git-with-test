package main

import "testing"

func TestOla(t *testing.T) {
	t.Run("diz olá para pessoas", func(t *testing.T) {
		resultado := Ola("Victor", "")
		esperado := "Olá, Victor"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Diz 'olá, mundo' quando string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Gabriella", "francês")
		esperado := "Bonjour, Gabriella"

		verificarMensagemCorreta(t, resultado, esperado)
	})
}

func verificarMensagemCorreta(t testing.TB, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%q', esperado '%q'", resultado, esperado)
	}
}
