package main

import "fmt"
import "math"

type titik struct {
	x, y float64
}
func main() {
	var t1, t2 titik 
	fmt.Scanln(&t1.x, &t1.y)
	fmt.Scanln(&t2.x, &t2.y)
	fmt.Println(akar((t2.x-t1.x)*(t2.x-t1.x)-(t2.y-t1.y)*(t2.y-t1.y)))
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}
