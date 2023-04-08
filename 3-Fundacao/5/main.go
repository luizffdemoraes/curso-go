package main

import (
	"fmt"
)

const a = "Hello, World!"

type ID int

var (
	b bool    = true     // false
	c int     = 10       // 0
	d string  = "Wesley" // vazio
	e float64 = 1.2      // +0.000000e+000
	f ID      = 1
)

func main() {
	// O que é um Array? Trabalha como uma variavel, tem um tamanho fixo e tipo fixo
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 8

	fmt.Println(len(meuArray) - 1)         // Pegar ultima posição do Array
	fmt.Println(meuArray[len(meuArray)-1]) // Pegar ultimo valor do Array
	fmt.Println(meuArray[0])               // Pegar primeiro valor do Array

	// Percorrer array
	// %d quando for digito e %v quando for string
	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d \n", i , v )
	}

}
