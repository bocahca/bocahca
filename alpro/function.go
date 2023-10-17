package main

import "fmt"

func main() {
	var bilangan int
	var result bool

	fmt.Scan(&bilangan)
	result = fgenap(bilangan)
	if result {
		fmt.Println("genap")
	}else {
		fmt.Println("ganjil")
	}
}

func fgenap(x int) bool {
	return x%2 == 0
}
