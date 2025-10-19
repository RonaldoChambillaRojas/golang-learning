package main

import "fmt"

func main() {

	// modulo := 5 % 2

	switch modulo := 4 % 2; modulo { // * FORMA DIRECTA
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// fmt.Println(modulo) // * CON LA FORMA DIRECTA YA NO PODEMO ACCEDER A LA VARIABLE MODULO. DESDE FUERA DE EL BLOQUE SE SWITCH

	// * SIN CONDICION:
	value := 45

	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}

}
