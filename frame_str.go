package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println("***")
	if n[0] == '-' {
		fmt.Println("*-*")
		n = n[1:]
	}
	length := len(n)
	for i := 0; i < length; i++ {
		fmt.Printf("*%c*\n", n[i])
	}
	fmt.Println("***")
}
