package main

import "fmt"
import "math"

func main() {
	var panjang, lebar, luas, keliling, diagonal float64
	fmt.Scanln(&panjang, &lebar)
	luas = panjang * lebar
	keliling = 2*panjang + 2*lebar
	diagonal = math.Sqrt(panjang*panjang + lebar*lebar)
	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)
	fmt.Println("Panjang diagonal:", diagonal)
}
