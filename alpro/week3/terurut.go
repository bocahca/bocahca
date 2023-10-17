package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	for x != y {
		urut(&x, &y)
		fmt.Println(x, y)
		fmt.Scanln(&x, &y)
	}
}
func urut(a, b *int) {
	var temp int
	if *a > *b {
		temp = *a
		*a = *b
		*b = temp
	}
}
