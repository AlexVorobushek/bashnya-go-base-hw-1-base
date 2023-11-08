package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	alreadyHave := make(map[string]bool)
	for _, str := range items {
		alreadyHave[str] = true
	}
	fmt.Println(len(alreadyHave))
}
