package main

import "fmt"

func main() {
	// * ARRAY
	// ? CONFIRMAR: los arrays en go no se pueden modificar su estructura.
	// ? pero si el valor del elementos que contiene.
	var array [4]int // ? OJO: si usamos esta version directamente tenemos un array con Zero values, en este caso array de 4 ceros.
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

}
