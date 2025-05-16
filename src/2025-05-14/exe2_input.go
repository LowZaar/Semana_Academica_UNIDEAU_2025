package main

import (
	"fmt"
)

func main() {
	var num int;

	fmt.Println("Digite um numero")
	fmt.Scan(&num)

	if num % 2 == 0 {
		fmt.Println(num, "é par")
	} else {
		fmt.Println(num, "é impar")
	}
}