package main

import "fmt" 
import "math"

type titik struct {
	x, y float64
}
	var t1, t2 titik
func main() {

	fmt.Scanln(&t1.x, &t1.y)
	fmt.Scanln(&t2.x, &t2.y)
	fmt.Println(jarak(t1,t2))

}

func jarak(p1, p2 titik) float64 {
	var jumlah float64
	jumlah = (p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y)
	return akar(jumlah)
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}