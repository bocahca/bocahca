package main

import "fmt"

type AsDos_T struct {
	kode, nama, mk string
	nim, jumlah int
}
const NMAX int = 2500 
type TabAsDos_T [NMAX] AsDos_T

func main() {
	var n int
	n = 0
	var matkul string
	var x TabAsDos_T
	bacaAsDos_1303220115(&x, &n)
	fmt.Scanln(&matkul)
	cetakAsdos_1303220115(x, n, matkul)
}

func bacaAsDos_1303220115(tabel *TabAsDos_T,  n *int) {
	var i int
	i = 0
	fmt.Scanln(&tabel[0].kode, &tabel[0].nama, &tabel[0].nim, &tabel[0].mk, &tabel[0].jumlah)
	for tabel[i].kode !="XXX" {
		i++
		*n++
		fmt.Scanln(&tabel[i].kode, &tabel[i].nama, &tabel[0].nim, &tabel[0].mk, &tabel[0].jumlah)
	}
}

func cetakAsdos_1303220115(tabel TabAsDos_T, n int, mk string) {
	var i int
	for i = 0; i <= n ; i++ {
		if tabel[i].mk == mk {
			fmt.Println(tabel[i])
		}
	}
}