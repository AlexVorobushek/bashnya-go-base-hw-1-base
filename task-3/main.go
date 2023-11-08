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

	res := make([]int, 1, n)
	res[0] = items[n-1]
	res = append(res, items...)

	items = nil

	for i := 0; i<n; i++ {
		fmt.Print(res[i], " ")
	}
}
