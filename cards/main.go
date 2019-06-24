package main

import (
	"fmt"
)


func main() {
	ints := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, int := range ints {
		if int % 2 == 0 {
			fmt.Printf("%v is Even\n", int)
		} else {
			fmt.Printf("%v is Odd\n", int)

		}
	}
}
