package main

import (
	"fmt"
)

func main() {
	var n, score_adr, score_pash int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var andrey, pasha int
		fmt.Scan(&andrey, &pasha)

		if andrey > pasha {
			score_adr++
		} else if andrey < pasha {
			score_pash++
		}
	}

	if score_adr > score_pash {
		fmt.Println("Andrey")
	} else if score_adr < score_pash {
		fmt.Println("Pasha")
	} else {
		fmt.Println("Draw")
	}
}
