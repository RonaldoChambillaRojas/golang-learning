package main

import "fmt"

func main() {
	//? como funciona make?
	m := make(map[string]int) //? CONFIRAMAR: las claves pueden ser de cualquier tipo?
	// b := map[string]int //? ERROR
	// b["valor1"] = 1
	// fmt.Println(b)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// * RECORRER UN MAP
	for i, v := range m {
		fmt.Println(i, v)
	}

	// * ENCONTRAR UN VALOR
	value, ok := m["Jose"] //? de donde sale el valor de ok?. en este caso true porque el valor existe.
	// value2 := m["Jose0"]
	// fmt.Println(value2) // * retorna su zero value. en este caso 0.

	value2, ok2 := m["Jose0"] // ? podemos capturar mas cosas aparte de que si el valor existe o no?

	fmt.Println(value2, ok2)

	fmt.Println(value, ok)

	// ? DUDA:
	// ? el profesor dijo que los maps son mas eficientes que manejar arrays o slices,
	// ? porque de forma nativa implementan "concurrencia", que es la concurrencia en-
	// ? este caso? y como nos favorese?
	// RPT: NO ES ASI

}
