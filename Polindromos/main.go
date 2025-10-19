package main

import "fmt"

func isPalindromo(texto string) {
	var textRevese string
	for i := len(texto) - 1; i >= 0; i-- {
		textRevese += string(texto[i])
	}

	if texto == textRevese {
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
