package main

import (
	"fmt"	
)

func main () {

	array := [6]int{4,8,15,16,23,42}
	elem := 15
	existe := elementoExiste(array, elem)

	if(existe) {
		fmt.Println(elem, "Existe")
	} else {
		fmt.Println(elem, "Não existe")
	}

}

func elementoExiste(array [6]int, elemento int) bool {

	for i := 0; i < len(array); i++ {

		if(array[i] == elemento) {
			return true
		}
	}

	return false
}