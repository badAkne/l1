package main

import "fmt"

func setBit(num int64, pos int64, bit bool) int64 {
	var res int64

	if bit {
		res = num | (1 << pos)
	} else {
		res = num &^ (1 << pos)
	}
	return res
}

func main() {
	var num, pos int64
	var bit bool
	fmt.Scan(&num, &pos)
	fmt.Scan(&bit)
	fmt.Println(setBit(num, pos, bit))
}
