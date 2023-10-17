package main

import "fmt"

func main() {
	var i, n_passed, n_failed, n int
	var n1, n2, n3, avg float64
	fmt.Println("Berapa jumlah siswa yang nilainya akan diproses?")
	fmt.Scanln(&n)
	n_passed = 0
	n_failed = 0 
	for i = 1; i<= n; i++{
		fmt.Scanln(&n1, &n2, &n3)
		avg = (n1+n2+n3) /3
		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed = n_passed + 1
		}else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed = n_failed + 1
		}
	}
	fmt.Println("Jumlah siswa lolos seleksi administrasi", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi administrasi", n_failed)
}