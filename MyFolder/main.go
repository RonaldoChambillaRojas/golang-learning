package main

import "fmt"

func main() {
	// * Declaracion de contantes.
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// * Declaracion de Varibles enteras.
	base := 12
	var altura int = 14
	var area int
	var otra = 2

	fmt.Println(base, altura, area, otra)

	// * Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// * Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area Cuadrado : ", areaCuadrado)

	// * Operadores

	x := 10
	y := 50

	// suma
	result := x + y
	fmt.Println("suma :", result)
	// resta
	result = y - x
	fmt.Println("resta :", result)
	// Multiplicacion
	result = y * x
	fmt.Println("multiplicacion :", result)
	// Divicion
	result = y / x
	fmt.Println("divicion :", result)
	// * Recordemos que la divicion tiene 4 partes: 1: dividendo, 2: divisor, 3: cosiente, 4: residuo | modulo

	// Modulo
	result = y % x
	fmt.Println("modulo: ", result)

	// incrementar
	x++
	fmt.Println(x)

	// decremental
	x--
	fmt.Println(x)
}
