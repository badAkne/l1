package main

import "fmt"

func pop[T any](slice []T, idx int) []T {
	if idx >= len(slice) || idx < 0 || slice == nil {
		return slice
	}

	copy(slice[idx:], slice[idx+1:])

	var empty T
	slice[len(slice)-1] = empty

	return slice[:len(slice)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	s := pop(arr, 3)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
