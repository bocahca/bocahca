package main

import (
	"fmt"
)
var a int
func main() {
	a = 10
	procB(&a, &a)
	fmt.Println(a)
}
func procB(b, c *int) {
	*b = *b+*c
	*c = a + *b + *c
}
