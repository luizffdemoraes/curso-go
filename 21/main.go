package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro)
	fmt.Println(carro.Andar())
	fmt.Printf("Resultado: %v \n", s)
	fmt.Println(matematica.A)

}
