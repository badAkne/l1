package main

import (
	"fmt"
	"math/rand"
)

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivotIdx := rand.Int() % len(arr)

	arr[pivotIdx], arr[right] = arr[right], arr[pivotIdx]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	qSort(arr[:left])
	qSort(arr[left+1:])

	return arr
}
func main() {
	arr := []int{7, 3, 2, 9, 10, 4, 52}

	fmt.Println(qSort(arr))
}
