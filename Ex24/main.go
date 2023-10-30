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
	inputNumbers := GetInputSlice()
	hours := inputNumbers[0]
	minutes := inputNumbers[1]
	var counter int
	fmt.Scan(&counter)

	for i := 0; i < counter; i++ {
		if hours >= 24 {
			hours = hours % 24
		}
		fmt.Printf("%02d:%02d\n", hours, minutes)

		minutes += 5
		hours += minutes / 60
		minutes %= 60
	}
}
