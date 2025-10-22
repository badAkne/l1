package main

import "fmt"

func foo(s1, s2 []int) []int {
	var res []int
	m := make(map[int]int)
	minLen := len(s1)

	if len(s2) < minLen {
		minLen = len(s2)
	}

	for i := 0; i < minLen; i++ {
		m[s1[i]]++
		m[s2[i]]++
	}

	for k, v := range m {
		if v != 1 {
			res = append(res, k)
		}
	}
	return res
}
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(foo(s1, s2))
}
