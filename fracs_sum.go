package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	n, m := uint64(1), uint64(1)
	nSum, mSum := 0.0, 0.0
	const max int = 248397

	for nSum < x {
		nSum += 1.0 / float64(n)
		n++
	}

	n--

	if nSum >= float64(x) {
		n--
	}

	for mSum < x {
		mSum += 1.0 / float64(m)
		m++
	}

	m--

	if mSum >= float64(x) {
		m--
	}

	fmt.Printf("%f %f\n", nSum, mSum)
	fmt.Printf("%d %d\n", n, m)
}
