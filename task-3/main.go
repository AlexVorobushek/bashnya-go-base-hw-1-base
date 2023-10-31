package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	items := make([]int, n)
	for i := 0; i<n; i++ {
		fmt.Scan(&items[i])
	}

	fmt.Print(items[n-1], " ")
	for i := 0; i<n-1; i++ {
		fmt.Print(items[i], " ")
	}
}