package main

import ("fmt")

const a = "Hello, World!"

//Criação de tipos
type ID int

// escopo global
var (
	b bool    = true     // false
	c int     = 10       // 0
	d string  = "Wesley" // vazio
	e float64 = 1.2      // +0.000000e+000
	f ID      = 1
)

// escopo local (Toda variavel declarada precisa ser utilizada)
// %T retorna o tipo
// %v retorna o valor
func main() {
	fmt.Printf("O valor de E é %v", e)
	fmt.Printf("O tipo de E é %T", e)
}
