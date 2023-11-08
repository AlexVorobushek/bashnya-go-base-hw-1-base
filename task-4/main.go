package main

import (
	"fmt"
)

func isSimmetry(items []int, n int) bool{
	for i:=0; i<n+1; i++{
		for j:=i+1; j < n; j++{
			if items[n*i+j]!=items[n*j+i]{
				return false
			}
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	items := make([]int, n*n)
	for i := 0; i<n*n; i++ {
		fmt.Scan(&items[i])
	}

	if isSimmetry(items, n){
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
