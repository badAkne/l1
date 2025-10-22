package main

import "fmt"

func zip(words []string) []string {
	res := []string{}
	m := make(map[string]int)

	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	for k := range m {
		res = append(res, k)
	}
	return res
}
func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(zip(words))
}
