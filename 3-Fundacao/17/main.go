package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c Cliente) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	//wesley := Cliente{
	//	nome: "Wesley",
	//}
	//wesley.andou()
	//fmt.Printf("O valor da struct com nome %v", wesley.nome)

	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
	println(NewConta())
}
