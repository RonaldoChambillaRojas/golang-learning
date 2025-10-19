package main

import "fmt"

func main() {

	// * BUCLE: for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------------------")

	// * BUCLE: for while

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// * BUCLE: forever
	// ! CUIDADO CON EL MAL USO DE ESTE
	// ! SIEMPRE ASEGURARSE DE QUE ALGO
	// ! PARE LA EJECUCION DE ESTA.

	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// * BUCLE: for range

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} // * HOJO: este es un tipo de dato Slice : lo mas similar a los arrays en js
	for indice, numPar := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", indice, numPar)
	}
}
