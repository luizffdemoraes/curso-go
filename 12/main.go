package main

import "fmt"

// linguagem go não é orientado há objetos
//o mais proximo é o struct

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco          // Composição
	// Address  Endereco Criando uma variavel do tipo
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome %s, Idade %d, Ativo %t", wesley.Nome, wesley.Idade, wesley.Ativo)

	wesley.Ativo = false
	wesley.Cidade = "São Paulo"

	fmt.Println(wesley)

}
