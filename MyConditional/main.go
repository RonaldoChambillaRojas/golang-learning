package main

import "fmt"

func main() {
	valor1 := 2
	// valor1 := "1" // ? CONFIRMAR: Porque me da error cuando trato de validar esto?.
	valor2 := 2

	// * CONDITIONALS

	if valor1 == 1 {
		fmt.Println("Es: 1")
	} else {
		fmt.Println("No es: 1")
	}

	// * COMPUERTAS LOGICAS
	// AND:

	// if valor1, valor2 == 2 {} // ? ERROR

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("AND: ambos valores cumplen la condicion")
	} else {
		fmt.Println("AND: uno o ambos valores no cumplen la condicion")
	}

	// OR:

	if valor1 == 1 || valor2 == 2 {
		fmt.Println("OR: uno o ambos valores cumplen la condicion")
	} else {
		fmt.Println("OR: ambos valores no cumplen la condicion")
	}

}
