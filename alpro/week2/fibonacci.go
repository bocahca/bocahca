package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fib(n))
}

func fib(bil int) int {
	var b1, b2, i int
	b1 = 0
	b2 = 1
	for i = 2; i <= bil; i++ {
		b1 = b2
		b2 = b1 + b2
	}
	return b1
}
