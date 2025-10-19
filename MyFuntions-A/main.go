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

func returnValue(a int) int {
	return a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument("numeros: ", 2, 3)
	value := returnValue(5)
	fmt.Println(value)
}
