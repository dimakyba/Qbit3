package main

import (
	"fmt"
)

func main() {
	var n, p, m float64
	fmt.Scan(&n, &p, &m)
	p /= 100.0
	i := 0
	for n < m {
		n = n*p + n
		i++
	}
	fmt.Println(i)
}
