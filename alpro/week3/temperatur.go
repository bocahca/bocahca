package main

import "fmt"

func main() {
	var C, R, F, K float64
	fmt.Scanln(&C)
	konversi(C, &R, &F, &K)
	fmt.Print(R,"R"," ", F,"F"," ",K,"K")
}
func konversi(celcius float64, reamur, farenheit, kelvin *float64) {
	*reamur = celcius * 4/5
	*farenheit = celcius * 9/5 +32
	*kelvin = 273.15 + celcius
}