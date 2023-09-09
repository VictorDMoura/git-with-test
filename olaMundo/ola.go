package main

import "fmt"

const (
	espanhol            = "espanhol"
	frances             = "francês"
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoOlaFrances   = "Bonjour, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {

	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
