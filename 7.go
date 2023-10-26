package main

import "fmt"

func main() {
	var i int
	k := 1
	counter := 0
	fmt.Scan(&i)
	for k != 0 {
		fmt.Scan(&k)
		if k > i {
			counter++
		}
		i = k
	}

	fmt.Println(counter)
}
