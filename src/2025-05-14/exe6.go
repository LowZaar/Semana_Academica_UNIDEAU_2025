package main

import (
	"fmt"
)

func main() {
	var num int
	var soma int = 1;

	fmt.Println("Digite o número: ")
	fmt.Scan(&num)


	for x := 1; x <= num; x++ {
		soma += x
	}

	fmt.Println("Soma final", soma)
}