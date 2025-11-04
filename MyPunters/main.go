package main

import (
	Pc "MyPunters/mipc"
	"fmt"
)

func main() {

	a := 50
	b := &a //? Si usamos el ampersand, asemos referencia a la ubicacion?
	c := a  //? sin el ampersand asemos referencia al valor?
	d := b

	fmt.Println(a)  //* Salida: 50
	fmt.Println(b)  //* Salida: 0xc0000120f8
	fmt.Println(*b) //?  con el asterisco la salida es 50. Osea lo que hacemos es apuntar al valor?.
	fmt.Println(c)  //* salida es: 50
	fmt.Println(*d) //* salida es: 50

	*b = 100
	fmt.Println(a) //* Aqui a recibe el valor de 100
	fmt.Println(c) //* Aqui c sigue siendo 50

	myPC := Pc.Pc{Ram: 16, Disk: 500, Brand: "msi"}
	fmt.Println(myPC)

	myPC.Ping()

	fmt.Println(myPC)

	myPC.DuplicateRAM()

	fmt.Println(myPC)

	myPC.DuplicateRAM()

	fmt.Println(myPC)

}
