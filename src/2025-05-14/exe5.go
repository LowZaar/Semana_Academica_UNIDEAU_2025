package main

import (
	"fmt"
)

func main() {
	
    for x := 1; x <= 100; x++ {
		var result string

		if x % 3 == 0 {
			result += "Foo"
		}

		if x % 5 == 0 {
			result += "Bar"
		}
		
		result = fmt.Sprintf("%d: %s", x, result)
		
		fmt.Println(result)
	}
}