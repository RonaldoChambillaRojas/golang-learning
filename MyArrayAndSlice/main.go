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

		goroutine 1 [running]: //? De donde sale este gorountine? aun no llegue a ese tema pero, de donde sale este? ya los use?
		main.main()
		        /home/tintodev/dev/curso-go/MyArrayAndSlice/main.go:25 +0xf8
		exit status 2
	*/

	fmt.Println(slice, len(slice), cap(slice))

	// * METODOS EN EL SLICE

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	fmt.Println("Antes del append:", len(slice), cap(slice))

	// * Append:
	slice = append(slice, 7)
	fmt.Println("append: ", slice)
	fmt.Println("Despues del append:", len(slice), cap(slice)) //? CONFIRMAR: aqui se duplica la capacidad.

	// * Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)       //? CONFIRMAR: esto se parese mucho al operador spread de js, solo que esta al reves.
	otroNombre := append(slice, newSlice...) //? CONFIRMAR: aqui supongo que estamos compartiendo el mismo array subyacente que slice.
	otroNombre[10] = 22                      //? CONFIRMAR: veo que aqui esto afecta el contenido de ambos slices.
	fmt.Println("otroNombre: ", otroNombre)
	fmt.Println("slice: ", slice)
	// extraSlice := append(slice, []int{11,12,13,14,15,16}) // ? CONFIRMAR: esto porque da error?.
	extraSlice := []int{11, 12, 13, 14, 15, 16}
	extraSlice = append(slice, extraSlice...)
	fmt.Println("extraSlice: ", extraSlice, len(extraSlice), cap(extraSlice))
	fmt.Println("---------------------")
	fmt.Println("otroNombre: ", otroNombre) //? CONFIRMAR: aqui ya no son modificados.
	fmt.Println("slice: ", slice)
	fmt.Println("verificar :", len(slice), cap(slice))

}
