package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("Enter the number of lines: ")
	fmt.Scan(&n)

	var str string

	fmt.Scan(&str)

	trimmed := strings.TrimSpace(str)
	for i := 1; i <= n; i++ {
		// numbers := input_line()
		for j := 0; j < len(trimmed); j++ {
			fmt.Printf("%d ", trimmed[j])
		}
		fmt.Println()
	}
}
