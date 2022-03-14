package main

import (
	"fmt"
)

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for _, elem2 := range a {
		elem2 = 100
		fmt.Println(elem2)

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [1 2 3 4 5]

}
