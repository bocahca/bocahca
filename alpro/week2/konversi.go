package main

import "fmt"

func main() {
	var jam, menit, detik int

	fmt.Scanln(&jam, &menit, &detik)
	fmt.Println("Hasil konversi:", kon(jam, menit, detik), "detik")
}
func kon(j, m, d int) int {
	return j*3600 + m*60 + d
}
