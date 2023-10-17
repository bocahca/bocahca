package main

import "fmt"

const pi float64 = 3.14

func main() {
	var r, t1, massa1, t2, massa2, hasil float64
	fmt.Scanln(&r)
	fmt.Scanln(&t1, &massa1)
	fmt.Scanln(&t2, &massa2)
	display(massa(volume(r, t1), massa1), massa(volume(r, t2), massa2), &hasil)
	if hasil == 0 {
		fmt.Println("BALANCE")
	} else if hasil > 0 {
		fmt.Println(hasil)
	} else {
		fmt.Println(-1 * hasil)
	}
}

func volume(l, t float64) float64 {
	return pi * l*l*t
}

func massa(v, mj float64) float64 {
	return v * mj
}

func display(kiri, kanan float64, hasil *float64) {
	*hasil = kiri - kanan
}
