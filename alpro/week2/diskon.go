package main

import "fmt"

func main() {
	var jumlah int
	var member bool
	var disc float64
	fmt.Scanln(&jumlah, &member)
	disc = float64(jumlah) * diskon(jumlah, member)
	fmt.Println(float64(jumlah) - disc)
}
func diskon(hrg int, mbr bool) float64 {
	if hrg > 100000 && mbr {
		return 0.1
	} else if hrg > 100000 && !mbr {
		return 0.05
	} else {
		return 0
	}
}
