package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage es para imprimir un mensage
func PrintMessage() {
	fmt.Println("Hola")
}
