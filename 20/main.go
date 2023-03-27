package main

type MyNumber int

// constraint
// Ao utilizar ~ antes do tipo habilite um tipo
// com a mesma tipagem
type Number interface {
	~int | float64
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Comparable Traz valores compativeis para comparação
// any e similar a uma interface vazia
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.20, "Joao": 2000.3, "Maria": 300.0}
	m3 := map[string]MyNumber{"Wesley": 1000, "Joao": 2000, "Maria": 3000}

	println(SomaInteiro(m))
	println(SomaFloat(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
