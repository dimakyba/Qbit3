package main

import (
	"fmt"
	// "math"
)

func main() {
	var i, k int
	counter := 0
	fmt.Scan(&i)

	for {
		fmt.Scan(&k)
		if k >= i {
			counter++
			i = k
		} else {
			break
		}
	}

	fmt.Println(counter)
}
