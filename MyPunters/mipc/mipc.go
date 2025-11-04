package mipc

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() { //? Aqui no entendi muy bien esta forma de declarar la funcion, con este: (myPC pc) porque se hace asi?.
	fmt.Println(myPC.Brand, "Pong")
}

func (myPC *Pc) DuplicateRAM() {
	myPC.Ram = myPC.Ram * 2
}
