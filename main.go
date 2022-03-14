package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	var t int
	var x [100]int
	var cnt int
	cnt = 0
	for i := 0; i < a; i++ {
		fmt.Scan(&t)
		x[i] = t
		if t > 0 {
			cnt++

		}
	}
	fmt.Printf("%d", cnt)

}
