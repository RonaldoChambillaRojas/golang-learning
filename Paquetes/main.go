package main

import (
	"fmt"
	"os"
)

func main() {

	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("HOME enviroment variable is noto set")
	} else {
		fmt.Printf("HOME enviroment variable: %s\n", envVar)
	}

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error creting file: %v\n", err)
		return
	}
	defer file.Close()
}
