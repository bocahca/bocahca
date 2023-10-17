package main

import "fmt"

const NMAX int = 12345
type prodi struct {
	nama, akreditasi string
	tahun, aktif, lulusan int
}

type fakultas struct {
	nama string
	arrProdi [NMAX] prodi
	N int
}
var fif fakultas
var j int

func main() {
	buatFakultas_1303220115(&fif)
	for j = 0; j <= 9; j++ {
		daftarProdi_1303220115(&fif)
	}
	fmt.Printf("Prodi %s memiliki mahasiswa dan lulusan terbanyak yaitu %d \n" , terbanyak_1303220115(fif).nama, terbanyak_1303220115(fif).aktif + terbanyak_1303220115(fif).lulusan)
	fmt.Printf("Prodi termuda adalah %s \n", fif.arrProdi[termuda_1303220115(fif)].nama)
	fmt.Printf("Akreditasi Prodi terbanyak di Fakultas informatika adalah %s \n", prestasi_1303220115(fif) )
	cetakProdi_1303220115(fif, prestasi_1303220115(fif))
}

func buatFakultas_1303220115(f *fakultas) {
	fmt.Scanln(&f.nama)
	f.N = 0
}

func daftarProdi_1303220115(f *fakultas) {
	fmt.Scanln(&f.arrProdi[f.N].nama, &f.arrProdi[f.N].akreditasi, &f.arrProdi[f.N].tahun, &f.arrProdi[f.N].aktif, &f.arrProdi[f.N].lulusan)
	if cekProdi_1303220115(f.arrProdi[f.N].nama, fif) == -1 {
		f.N++
	}else {
		f.arrProdi[cekProdi_1303220115(f.arrProdi[f.N].nama, fif)].aktif += f.arrProdi[f.N].aktif
		f.arrProdi[cekProdi_1303220115(f.arrProdi[f.N].nama, fif)].lulusan += f.arrProdi[f.N].lulusan
	}

}

func cekProdi_1303220115(s string, f fakultas) int {
	var j int
	j = 0
	for j < f.N &&  f.arrProdi[j].nama !=s {
		j++
	}
	if j == f.N {
		return -1
	}else {
		return j
	}
}

func terbanyak_1303220115(f fakultas) prodi {
	var idxmax int
	idxmax = 0
	for i:= 1; i < f.N; i++ {
		if f.arrProdi[i].aktif + f.arrProdi[i].lulusan > f.arrProdi[idxmax].aktif + f.arrProdi[idxmax].lulusan {
			idxmax = i
		}
		
	}
	return f.arrProdi[idxmax]
}

func termuda_1303220115(f fakultas) int {
	var idxmax int
	idxmax = 0
	for i:= 1; i < f.N; i++ {
		if f.arrProdi[i].tahun >= f.arrProdi[idxmax].tahun {
			idxmax = i
		}
	}
	return idxmax
}

func prestasi_1303220115(f fakultas) string {
	var jmlunggul, jmlbaik, jmlcukup, i int
	var terbanyak string
	for i = 0; i < f.N; i++ {
		if f.arrProdi[i].akreditasi == "unggul" {
			jmlunggul++
		}else if f.arrProdi[i].akreditasi == "baik" {
			jmlbaik++
		}else if f.arrProdi[i].akreditasi == "cukup" {
			jmlcukup++
		}
		terbanyak = "unggul"
		if jmlbaik > jmlunggul {
			terbanyak = "baik" 
			if jmlcukup > jmlbaik {
				terbanyak = "cukup"
			}
		}else if jmlcukup > jmlunggul {
			terbanyak = "cukup"
		}
	}
	return terbanyak
}

func cetakProdi_1303220115(f fakultas, x string) {
	for i:=0; i < f.N; i++ {
		if f.arrProdi[i].akreditasi == x {
			fmt.Print(f.arrProdi[i].nama, " ")
		}
	}
}