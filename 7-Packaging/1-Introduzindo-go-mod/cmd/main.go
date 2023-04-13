package main

import (
	"fmt"
	"github.com/luizffdemoraes/curso-go/7-Packaging/1-Introduzindo-go-mod/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
