package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func GetInputSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return numbers(scanner.Text())
}

func main() {
	var n, score_adr, score_pash int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		inputNumbers := GetInputSlice()
		andrey := inputNumbers[0]
		pasha := inputNumbers[1]
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
