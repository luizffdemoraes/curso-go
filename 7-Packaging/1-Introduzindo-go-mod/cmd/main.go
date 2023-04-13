package main

import (
	"fmt"
	"github.com/luizffdemoraes/curso-go/7-Packaging/1-Introduzindo-go-mod/math"
)

func main() {
	m := math.NewMath(1, 2)
	// m := math.Math{}
	m.C = 3
	fmt.Println(m.C)
	fmt.Println(m.Add())
	fmt.Println(math.X)
}
