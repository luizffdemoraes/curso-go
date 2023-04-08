package main

import "fmt"

// linguagem go não é orientado há objetos
//o mais proximo é o struct

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome %s, Idade %d, Ativo %t", wesley.Nome, wesley.Idade, wesley.Ativo)

	wesley.Ativo = false
	fmt.Println(wesley)

}
