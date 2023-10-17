package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}

func f(x int) int {
	// mengembalikan fungsi f(x)
	return x * x
}

func g(x int) int {
	// mengembalikan fungsi g(x)
	return x - 2
}

func h(x int) int {
	// mengembalikan fungsi h(x)
	return x + 1
}
