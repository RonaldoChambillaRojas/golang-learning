package main

import (
	"fmt"
	"log"
	"strconv"
)

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

	// * COMVERTIR TEXTO A NUMERO:

	value, err := strconv.Atoi("texto")
	// if err {    // ? OJO: no podemos validar de esta forma si existe o no el error
	// 	log.Fatal(err)
	// }

	if err != nil {
		log.Fatal(err) // ? ACLARAR: este log.Fatal como funciona? veo que para la ejecucion y lansa un error. fmt no puede hacer lo mismo?
	}
	// * err:
	/*
		No es: 1
		AND: uno o ambos valores no cumplen la condicion
		OR: uno o ambos valores cumplen la condicion
		2025/10/19 13:37:47 strconv.Atoi: parsing "texto": invalid syntax
		exit status 1
	*/

	fmt.Println("Value:", value)

}
