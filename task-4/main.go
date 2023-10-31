package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	items := make([]int, n*n)
	for i := 0; i<n*n; i++ {
		fmt.Scan(&items[i])
	}

	for i:=0; i<n+1; i++{
		for j:=i+1; j < n; j++{
			if items[n*i+j]!=items[n*j+i]{
				fmt.Println("no")
				os.Exit(0)
			}
		}
	}

	fmt.Println("yes")
}