package main

import "fmt"

func main() {
	var panjang, lebar, luas int
	fmt.Scanln(&panjang, &lebar)
	luas = panjang * lebar
	fmt.Println(luas)
}
