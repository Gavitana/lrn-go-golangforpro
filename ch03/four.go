package main

import "fmt"

type Power4 int

const (
	p4_0 Power4 = 1 << iota
	_
	p4_1
	_
	p4_2
	_
	p4_3
	_
	p4_4
	_
	p4_5
)

func main() {
	fmt.Println("4^0:", p4_0)
	fmt.Println("4^1:", p4_1)
	fmt.Println("4^2:", p4_2)
	fmt.Println("4^3:", p4_3)
	fmt.Println("4^4:", p4_4)
	fmt.Println("4^5:", p4_5)
}
