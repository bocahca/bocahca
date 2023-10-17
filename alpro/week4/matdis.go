package main

import "fmt"
var faktorial int
func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Println(permutasi(a,c), kombinasi(a,c))
	fmt.Println(permutasi(b,d), kombinasi(b,d))


}

func mencariFaktorial(n int, faktorial *int) {
	var i int
	*faktorial = 1
	for i = 1; i<= n ; i++{
		*faktorial = *faktorial * i
	}
}

func permutasi(n, r int) int {
	var a, b int
	mencariFaktorial(n, &faktorial)
	a = faktorial
	mencariFaktorial(n-r, &faktorial)
	b = faktorial
	return a/b
}

func kombinasi(n,r int) int {
	var a,b,c int
	mencariFaktorial(n, &faktorial)
	a = faktorial
	mencariFaktorial(r, &faktorial)
	b = faktorial
	mencariFaktorial(n-r, &faktorial)
	c = faktorial
	return a/(b*c)
}

