package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10
	// Exibir valor do ponteiro
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	b := &a
	*b = 30
	println(b)
	println(*b)
}
