package main

import (
	"fmt"
)

func main() {

	array := [5]int{5,10,22,53,31}
	var soma int

	for x := 0; x < len(array); x++ {
		soma += array[x]
	}

	fmt.Println("Array:", array)
	fmt.Println("Soma total:", soma)
}