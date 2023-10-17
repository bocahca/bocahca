package main

import "fmt"

func main() {
	var panjang, lebar, skala int
	var inP, inL int
	var outP, outL float64
	var s string
	fmt.Scanln(&panjang, &lebar)
	fmt.Scanln(&s, &skala)
	if s == "+" {
		zoomIn(panjang, lebar, skala, &inP, &inL)
		fmt.Println(inP, inL)
	} else if s == "-" {
		zoomOut(panjang, lebar, skala, &outP, &outL)
		fmt.Println(outP, outL)
	}
}

func zoomIn(p, l, s int, hasilP, hasilL *int) {
	*hasilP = p * s
	*hasilL = l * s
}

func zoomOut(p, l, s int, hasilP, hasilL *float64) {
	*hasilP = float64(p) * (1 / float64(s))
	*hasilL = float64(l) * (1 / float64(s))
}
