package main

import (
	"fmt"
	"strings"
)

func isPalindromo(texto string) {

	// ? CONFIRMAR:
	// ? esto lo copie de un compañero, algo que voy notando esque a diferencia de js-
	// ? que tenemos las propidades y metodos dentro de cada tipo de dato, aqui en go-
	// ? me da la secacion de que estos metodos como que estan separados de los tipos-
	// ? de datos, o algo asi.
	textoMin := strings.ToLower(texto)

	var textRevese string
	for i := len(textoMin) - 1; i >= 0; i-- {

		// ? CONFIRMAR:
		//? ??: Esta parte me parece muy interesante, se PARESE a la concatenacion de javascript-
		// ? cuando sumamos dos cadenas de texto, pero a la ves no da la sensacion de que sea lo mismo.
		// ? QUE PASA AQUI??
		textRevese += string(textoMin[i])
	}

	if textoMin == textRevese {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}
func main() {
	slice := []string{"Hola", "que", "haces"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("amor a roma")
	isPalindromo("casas")
}
