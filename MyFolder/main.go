package main

import "fmt"

func main() {
	// Declaracion de contantes.
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de Varibles enteras.
	base := 12
	var altura int = 14
	var area int
	var otra = 2

	fmt.Println(base, altura, area, otra)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
