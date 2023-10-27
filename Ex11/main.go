package main

import (
	"fmt"
	"math"
)

func main() {
	var current_max, prev_max int
	prev_max = math.MaxInt
	counter := 0
	current_max = 1

	for current_max != 0 {
		fmt.Scan(&current_max)
		if current_max == prev_max {
			counter++
		} else {
			counter = 0
		}

		prev_max = current_max
	}

	fmt.Println(counter)
}
