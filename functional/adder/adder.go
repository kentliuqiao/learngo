package main

import (
	"fmt"
)

// 闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Println(s)
	}
}
