package main

import (
	"fmt"
)

func main() {
	// Declarando um Slice -> Considerado um array infinito
	// Três partes um pontero, um tamanho e capacidade
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Zerar a capacidade e não retornar nada
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Retornar só os 4 primeiros itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Ignora os dois primeiros itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Aumentar a capacidade do slice
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
