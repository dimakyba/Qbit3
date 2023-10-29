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
	total_minutes := hours*60 + minutes
	var counter int
	fmt.Scan(&counter)
	for i := 1; i <= counter; i++ {
		if hours%24 != 0 {
			hours %= 24
		}
		fmt.Printf("%02d:%02d\n", hours, minutes)
		total_minutes += 5
		hours = total_minutes / 60
		minutes = total_minutes % 60

	}
}
