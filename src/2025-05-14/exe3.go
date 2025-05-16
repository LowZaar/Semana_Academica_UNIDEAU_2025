package main

import (
	"fmt"
)

func main() {
	var nome string

	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)

	//fmt.Println("Olá", nome," como você está?")
	result := fmt.Sprintf("Olá %s, como você está?", nome)

	fmt.Println(result)
}