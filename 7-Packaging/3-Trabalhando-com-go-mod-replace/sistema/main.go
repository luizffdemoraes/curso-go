package main

import "github.com/luizffdemoraes/curso-go/7-Packaging/3/math"

func main() {
	m := math.NewMath(1,2)
	println(m.Add())
}