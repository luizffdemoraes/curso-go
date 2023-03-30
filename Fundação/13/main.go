package main

import "fmt"

// linguagem go não é orientado há objetos
//o mais proximo é o struct
// Go Way

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Composição
	// Address  Endereco Criando uma variavel do tipo
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Ativo = false
	wesley.Desativar()
}
