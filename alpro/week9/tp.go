package main

import (
"math"
"fmt"
)

type poin struct {
 	x float64
    y float64
}

func main() {
	var A, B, C, D poin
	fmt.Scanln(&A.x, &A.y, &B.x, &B.y)
	fmt.Scanln(&C.x, &C.y, &D.x, &D.y)
	if jarak_1303220115(A, B) < jarak_1303220115(C, D) {
		fmt.Printf("Titik terdekat adalah titik A(%g,%g) dengan B(%g,%g) dengan jarak %.1f",A.x, A.y, B.x, B.y, jarak_1303220115(A,B))
		
	}else {
		fmt.Printf("Titik terdekat adalah titik C(%g,%g) dengan D(%g,%g) dengan jarak %.1f",C.x, C.y, D.x, D.y, jarak_1303220115(C,D))
	}
}

func jarak_1303220115(p1, p2 poin) float64 {
	var d float64
	d = (p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y)
	return math.Sqrt(d)
}


