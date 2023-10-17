package main

import "fmt"

func main() {
	var tahun,  A, CD int
	var diskon, hrgdisc, total float64
	fmt.Scanln(&tahun)
	fmt.Scanln(&total)
	A = tahun / 1000
	CD = tahun % 100
	diskon = float64(A) * float64(CD) / 100
	hrgdisc = total * diskon
	fmt.Println("besar diskon: ", diskon*100, "%")
	fmt.Println("Jumlah yang dibayar: ", total-hrgdisc)
}
