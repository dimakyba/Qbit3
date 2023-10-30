ckage main

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
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		inputNumbers := numbers(line)

		if len(inputNumbers) != 2 {
			return
		}

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
