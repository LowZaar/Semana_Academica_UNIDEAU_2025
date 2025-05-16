package main

import (
	"fmt"
	"math"
)

func main() {

	array := [6]int{4,8,15,16,23,42}
	var maior = 0
	var menor = math.MaxInt

	for x := 0; x < len(array); x++ {
		if array[x] > maior {
			maior = array[x]
		}

		if array[x] < menor {
			menor = array[x]
		}
	}

	fmt.Println("Array:", array)
	fmt.Println("Maior:", maior)
	fmt.Println("Menor:", menor)
}