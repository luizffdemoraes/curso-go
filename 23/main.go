package main

//  go build main.go
//  go build
//  go build -o xpto
func main() {
	a := 1
	// b := 2
	// c := 3

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

	// e && ou ||
	// if a > b && c > a {
	// 	println("a > b && c > a")
	// }

	// if a > b {
	// 	println(a)
	// } else {
	// 	println(b)
	// }
}
