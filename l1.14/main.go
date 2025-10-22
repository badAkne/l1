package main

import (
	"fmt"
	"reflect"
)

func identify(value interface{}) {
	/*
		switch value.(type) {
		case int:
			fmt.Println("You entered an int!")
		case string:
			fmt.Println("You entered an string!")
		case bool:
			fmt.Println("You entered an boolean!")
		case chan :
			fmt.Println("You entered a channel!")
		}
	*/

	fmt.Println(reflect.TypeOf(value))
}
func main() {
	ch := make(chan any)
	num := 1
	s := "yes"
	b := true
	identify(ch)
	identify(num)
	identify(s)
	identify(b)
}
