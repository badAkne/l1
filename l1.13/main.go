package main

import "fmt"

func swap(a, b int) (int, int) {
	a = a + b
	b = b - a
	b = -b
	a = a - b
	return a, b
}
func main() {
	a := 10
	b := 13
	fmt.Printf("before swap\na=%d\nb=%d\n", a, b)

	a, b = swap(a, b)

	fmt.Printf("after swap\na=%d\nb=%d\n", a, b)
}
