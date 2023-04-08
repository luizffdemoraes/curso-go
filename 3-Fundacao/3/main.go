package main

const a = "Hello, World!"
//Criação de tipos
type ID int

// escopo global
var (
	b bool    = true     // false
	c int     = 10       // 0
	d string  = "Wesley" // vazio
	e float64 = 1.2      // +0.000000e+000
	f ID 	  = 1
)

// escopo local (Toda variavel declarada precisa ser utilizada)
func main() {
	println(f)
}
