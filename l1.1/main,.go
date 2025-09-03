package main

import "fmt"

type Human struct {
	name   string
	age    int
	height int
}

type Action struct {
	Human
}

func (h *Human) greet() {
	fmt.Printf("Я чувак, меня зовут %s и мне %d лет!\n", h.name, h.age)
}

func (p *Human) move() {
	fmt.Printf("Я чувак, и я иду кое-куда!\n")
}

func (p *Human) listen() {
	fmt.Printf("Я чувак и ща слушаю репчик!")
}

func main() {
	h := Human{
		name:   "Стас",
		age:    19,
		height: 180,
	}
	a := Action{}

	a.greet()
	h.greet()
}
