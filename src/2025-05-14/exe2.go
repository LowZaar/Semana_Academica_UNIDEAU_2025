package main

import (
	"fmt"
)

func main() {
	var num int = 5

	if num % 2 == 0 {
		fmt.Println(num, "é par")
	} else {
		fmt.Println(num, "é impar")
	}

}