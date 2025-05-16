package main

import (
	"fmt"
)

func main() {

	var array [5]int

	for x := 0; x < len(array); x++ {
		array[x] = x + 1
	}

	fmt.Println("Array:", array)
}