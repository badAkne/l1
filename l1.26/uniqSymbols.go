package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	runeMap := make(map[rune]int)
	s = strings.ToLower(s)

	for _, char := range s {
		if _, exist := runeMap[char]; exist {
			return false
		}

		runeMap[char]++
	}

	return true
}

func main() {
	s := "abcd"
	s2 := "abCdefAaf"
	s3 := "abc😡😡"
	s4 := "abc😡!@##"
	s5 := "abc😡!@#"

	fmt.Println(isUnique(s))
	fmt.Println(isUnique(s2))
	fmt.Println(isUnique(s3))
	fmt.Println(isUnique(s4))
	fmt.Println(isUnique(s5))

}
