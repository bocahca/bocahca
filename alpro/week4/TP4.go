package main

import "fmt"

func main() {
	var tgl1, bln1, thn1 int
	var tgl2, bln2, thn2 int
	
	inputTglPinjam_1393220115(&tgl1, &bln1, &thn1)
	for valid_1303220115(tgl1, bln1, thn1) {
		hitungTanggalKembali_1303220115(tgl1, bln1, thn1, &tgl2, &bln2, &thn2)
		fmt.Println(tgl2, bln2, thn2)
		inputTglPinjam_1303220115(&tgl1, &bln1, &thn1)
	}
	if !valid_1303220115(tgl1, bln1, thn1) {
		fmt.Println("input tidak valid")
	}
}

func inputTglPinjam_1303220115(tanggal, bulan, tahun *int) {
	fmt.Scanln(&*tanggal, &*bulan, &*tahun)
}
func valid_1303220115(tanggal, bulan, tahun int) bool {
	var jmlHari int
	getJumlahHari_1303220115(bulan, tahun, &jmlHari)
	return (tahun > 0) && (1 <= bulan && bulan <= 12) && (0 < tanggal && tanggal <= jmlHari)
}

func kabisat_1303220115(tahun int) bool {
	return (tahun%400 != 0 && tahun%100 != 0 && tahun%4 == 0) || tahun%400 == 0
}

func getJumlahHari_1303220115(bulan, tahun int, jmlHari *int) {
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		*jmlHari = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		*jmlHari = 30
	} else if kabisat_1303220115(tahun) && bulan == 2 {
		*jmlHari = 29
	} else {
		*jmlHari = 28
	}
}

func hitungTanggalKembali_1303220115(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
	var jmlHari int
	getJumlahHari_1303220115(bulan1, tahun1, &jmlHari)
	if valid_1303220115(tanggal1, bulan1, tahun1) {
		if bulan1 == 12 && tanggal1 + 3 > jmlHari {
			*tanggal2 = tanggal1 + 3 - jmlHari
			*bulan2 = 1
			*tahun2 = tahun1 + 1
		}else if tanggal1+3 > jmlHari {
			*tanggal2 = tanggal1 + 3 - jmlHari
			*bulan2 = bulan1 + 1
			*tahun2 = tahun1
		} else {
			*tanggal2 = tanggal1 + 3
			*bulan2 = bulan1
			*tahun2 = tahun1
		}
	}
}
