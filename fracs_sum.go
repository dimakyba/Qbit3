package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n big.Float
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// First loop
	N := new(big.Float)
	i := big.NewFloat(1.0)
	one := big.NewFloat(1.0)
	for {
		fraction := new(big.Float).Quo(one, i)
		N = N.Add(N, fraction)

		cmp := n.Cmp(N)
		if cmp <= 0 {
			result := new(big.Float).Sub(i, one)
			fmt.Printf("%.0f ", result)
			break
		}

		i.Add(i, one)
	}

	// Reset N
	N = new(big.Float)

	// Second loop
	i = big.NewFloat(1.0)
	for {
		fraction := new(big.Float).Quo(one, i)
		N = N.Add(N, fraction)

		cmp := n.Cmp(N)
		if cmp < 0 {
			fmt.Printf("%.0f\n", i)
			break
		}

		i.Add(i, one)
	}
}
