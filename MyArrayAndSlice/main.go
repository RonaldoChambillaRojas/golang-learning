package main

import "fmt"

func main() {
	// * ARRAY
	// ? CONFIRMAR: los arrays en go no se pueden modificar su estructura.
	// ? pero si el valor del elementos que contiene.
	var array [4]int // ? OJO: si usamos esta version directamente tenemos un array con Zero values, en este caso array de 4 ceros.

	// ? CONFIRMAR: esto parece ser valido, y parece indicar que podemos agregar menos o la misma cantida de valores espeficados mas no mas.
	// otroArray := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Otro Array:", otroArray)
	array[0] = 1
	array[1] = 2
	// array[5] = 4 //! ERROR

	// ? CONFIRMAR: en este caso cual es la diferencia entre len y cap? parese que hacen lo mismo.
	// ? cuando debo usar uno u otro?
	fmt.Println(array, len(array), cap(array))

	// * SLICE
	// ? CONFIRMAR: lo unico que cambia es que no le dicimos la longitud del array?
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	// sliceV2 := []int
	// sliceV2[0] = 5  	// ? PORQUE ESTO DA ERROR?

	// slice[8] = 10 // ? !!!!OJO: aqui go no me da erro de sintaxis. pero me da error al ejecutar.
	// * ERROR:
	/*
			[1 2 0 0] 4 4
		panic: runtime error: index out of range [8] with length 7

		goroutine 1 [running]:
		main.main()
		        /home/tintodev/dev/curso-go/MyArrayAndSlice/main.go:25 +0xf8
		exit status 2
	*/

	fmt.Println(slice, len(slice), cap(slice))

	// * METODOS EN EL SLICE
	// ? CONFIRMAR: en esta parte siento que el profe no se expreso de una buena manera,
	// ? mesclo indices con valores en su explicion y no se dio a enteder bien.
	// ? pero aqui solo hablamos de indices verdad?, nada que ver los valores.
	// ? o talves si?

	fmt.Println(slice[0])
	fmt.Println(slice[:3])  // ? aqui nos trae todo incluyendo el valor en el posicion 3.
	fmt.Println(slice[2:4]) // ? por eso en este caso le decimos: trame a partir de la posicion 2 hasta la posicion 4. por eso no incluy al 1.
	fmt.Println(slice[4:])  // ? aca igual, no incluye el valor 3, solo nos trae los siguientes.

	// * Append:
	slice = append(slice, 7)
	fmt.Println(slice)

	// * Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) //? CONFIRMAR: esto se parese mucho al operador spread de js, solo que esta al reves.
	//otroNombre = append(slice, newSlice...) //? CONFIRMAR: esto da error. no es que estemos creando un nuevo slice. que estamos haciendo?
	fmt.Println(slice)

}
