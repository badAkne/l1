package main

import (
	"fmt"
	"strings"
)

func reverseSentence(s string) string {
	sent := strings.Split(s, " ")

	for i, j := 0, len(sent)-1; i < j; i, j = i+1, j-1 {
		sent[i], sent[j] = sent[j], sent[i]
	}

	return strings.Join(sent, " ")
}

func main() {
	s := ""
	fmt.Println(reverseSentence(s))
}
