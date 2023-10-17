package main

import "fmt"

func main() {
	var x, y, permutasi int

	fmt.Scanln(&x, &y)
	if x >= y {
		permutasi = fak1(x) / fak2(x, y)
	} else {
		permutasi = fak1(y) / fak2(y,x)
	}
	fmt.Println(fak1(x), fak1(y), permutasi)

}
func fak1(b1 int) int {
	var total, i int
	total = 1
	for i = 1; i <= b1; i++ {
		total = total * i
	}
	return total
}
func fak2(b1, b2 int) int {
	var result, j, jml int
	result = 1
	jml = b1 - b2
	for j = 1; j <= jml; j++ {
		result = result * j
	}
	return result
}
