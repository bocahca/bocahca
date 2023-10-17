package main

import (
	 "fmt"
	 "math"
)

func main() {
	var a,b int
	fmt.Scanln(&a, &b)
	fmt.Printf("%.3f",z_1303220115(a,b))
	fmt.Print(" ")
	fmt.Printf("%.3f", z_1303220115(b,a)) 
}

func z_1303220115(x,y int) float64 {
	return math.Sqrt((6*float64(y)*3*float64(x)) / (4*float64(y)))
}

