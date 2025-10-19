package main

import "fmt"

func main() {
	// * DEFER
	// ? CONFIRMAR: cuando usamos defer en una funcion lo que hace es asta el final.
	// ? pero el final de que? sera de su funcion padre o algo asi? o sera talves
	// ? del blocke donde se encuentre?, solo funciona para aplicarselas a funciones/methodos?
	// ? que pasa si tengo mas de un defer?
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// * CONTINUE && BREAK
	// ? CONFIRMAR: Estos keywords solo podemos usarlos dentro de bucles for?.
	for i := 0; i < 10; i++ {

		fmt.Println(i) // ? OJO: Si ejecutamos esto aqui se imprime el numero 2

		// Continue:
		// ? CONFIRMAR: por lo que vi cuando usamos continue, lo que hace es parar la
		// ? iteracion en la que se encuentra y saltarse a la siguiente, omitiendo el
		// ? resto de codigo que sigue.
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// fmt.Println(i) // ? OJO: Aqui ya no se imprime el numero 2.

		// Break:
		// ? CONFIRMAE: por lo que vi, simplemente para el bucle completo.
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
