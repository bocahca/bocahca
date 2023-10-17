package main

import "fmt"

func main() {
	var b1, b2, b3 int
	fmt.Scanln(&b1, &b2, &b3)
	if segitiga(b1, b2, b3) {
		fmt.Println("segitiga")
	}else {
		fmt.Println("bukan segitiga")
	}
}
func segitiga(a, b, c int) bool {
	return a+b > c && a + c > b && b + c > a && a!= 0 && b != 0 && c != 0
}