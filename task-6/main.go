package main

import (
	"fmt"
	"strings"
	"bufio"
  	"os"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
	items := strings.Split(text, " ")

	var already_have []string
	var ans int
	for _, str := range items{
		if ! stringInSlice(str, already_have) {
			ans++
			already_have = append(already_have, str)
		}
	}
	fmt.Println(ans)
}