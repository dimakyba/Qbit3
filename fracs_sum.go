package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Scan(&n)

	sum := 0.0
	i := 1.0

	for sum < n {
		sum += 1 / i
		i++
	}
	fmt.Printf("%.0f ", i-2)
	sum = 0.0
	i = 1.0

	for sum <= n {
		sum += 1 / i
		i++
	}
	fmt.Printf("%.0f\n", i-1)
}
