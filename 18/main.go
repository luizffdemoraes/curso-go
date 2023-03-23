package main

import "fmt"

/*
-- Todos que tiverem o método y vão corresponder ao tipo x.
-- Caso não tenha nenhum método, significa que a interface implementa todos.
-- Utilize com moderação.
type x interface {}
*/

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v \n", t, t)
}
