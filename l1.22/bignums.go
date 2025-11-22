package main

import (
	"fmt"
	"math/big"
)

func Sum(x, y *big.Int) *big.Int { return x.Add(x, y) }

func Substract(x, y *big.Int) *big.Int { return x.Sub(x, y) }

func Div(x, y *big.Int) *big.Int {

	if y.Sign() == 0 {
		return big.NewInt(0)
	}

	return x.Div(x, y)
}

func Multiply(x, y *big.Int) *big.Int { return x.Mul(x, y) }

func main() {
	x, ok := big.NewInt(0).SetString("03092384029348203840239840754648763458734954395864395", 10)
	if !ok {
		panic("can't create x")
	}

	y, ok := big.NewInt(0).SetString("28734983274928345645674", 10)
	if !ok {
		panic("can't create y")
	}

	fmt.Println(Sum(x, y))

	fmt.Println(Substract(x, y))

	fmt.Println(Div(x, y))

	fmt.Println(Multiply(x, y))
}
