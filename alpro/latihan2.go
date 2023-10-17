package main

import "fmt"

func main() {
	var b1, b2, b3, b4 int
	var h1, h2, h bool
	fmt.Scanln(&b1, &b2, &b3, &b4)
	h1 = b1 != b2
	h2 = b3 <= b4
	h = !h1 && !h2
	fmt.Println("Hasil:", h)
}	