package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

// * HOJO: aqui podemos declarar el tipo de la varibl por separado: b int, c int
// * o junto: b,c int
// * SIMPRE Y CUANDO SEAN DEL MISMO TIPO OBVIAMENTE.
func tripleArgument(a string, b, c int) {
	fmt.Println(a, b, c)
}

// ! HOJO: Es obligatorio decirle a go el tipo de dato de retorno.
// * Forma para retornar un solo valor.
func returnValue(a int) int {
	return a * 2
}

// ? CONFIRMAR: Aqui el profesor especifico el tipo de salida de esta manera: (b, c int)
// ? pero le puse asi y funcina igual.
func dobleReturn(a int) (int, int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument("numeros: ", 2, 3)
	value := returnValue(5)
	fmt.Println(value)
	// value1, value2 := dobleReturn(10) // * HOJO: esta forma es para usar ambos retornos
	// fmt.Println("value1:", value1, "value2:", value2)

	value1, _ := dobleReturn(10)   // * HOJO: esta forma es para descartar algun retorno
	fmt.Println("value1:", value1) // * HOJO: aqui ya no podemos usar _ para represar un valor de retorno.
}
