package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	for _, ch := range s{
		if ch=='1'{
			fmt.Print("one")
		} else{
			fmt.Printf("%c", ch)
		}
	}
}