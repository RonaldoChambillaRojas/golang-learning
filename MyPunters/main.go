package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() { //? Aqui no entendi muy bien esta forma de declarar la funcion, con este: (myPC pc) porque se hace asi?.
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a //? Si usamos el ampersand, asemos referencia a la ubicacion?
	c := a  //? sin el ampersand asemos referencia al valor?
	d := b

	fmt.Println(a)  //* Salida: 50
	fmt.Println(b)  //* Salida: 0xc0000120f8
	fmt.Println(*b) //?  con el asterisco la salida es 50. Osea lo que hacemos es apuntar al valor?.
	fmt.Println(c)  //* salida es: 50
	fmt.Println(*d)

	*b = 100
	fmt.Println(a) //* Aqui a recibe el valor de 100
	fmt.Println(c) //* Aqui c sigue siendo 50

	myPC := pc{ram: 16, disk: 500, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)

	myPC.duplicateRAM()

	fmt.Println(myPC)

	myPC.duplicateRAM()

	fmt.Println(myPC)

}
