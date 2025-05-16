package main

import (
	"fmt"
)

func main() {
	var numero1 int
	var numero2 int

	fmt.Println("Digite o primeiro número: ")
	fmt.Scan(&numero1)

	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&numero2)

	adicao := fmt.Sprintf("%d + %d = %d", numero1, numero2, numero1 + numero2)
	fmt.Println("Adição:", adicao)

	subtracao := fmt.Sprintf("%d - %d = %d", numero1, numero2, numero1 - numero2)
	fmt.Println("Subtração:", subtracao)

	multiplicacao := fmt.Sprintf("%d * %d = %d", numero1, numero2, numero1 * numero2)
	fmt.Println("Multiplicação:", multiplicacao)

	divisao := fmt.Sprintf("%d / %d = %d", numero1, numero2, numero1 / numero2)
	fmt.Println("Divisão", divisao)


}