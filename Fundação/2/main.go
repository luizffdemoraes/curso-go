package main

const a = "Hello, World!"

// escopo global
var (
	b bool    = true     // false
	c int     = 10       // 0
	d string  = "Wesley" // vazio
	e float64 = 1.2      // +0.000000e+000
)

// escopo local (Toda variavel declarada precisa ser utilizada)
func main() {
	a := "X" // É feito da primeira vez := quando a variavel é criada
	a = "XX"
	println(a)
}

func x() {

}
