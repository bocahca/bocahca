package main

import "fmt"

func main() {
	var modal, totalUang, totalKeluar, sisa, keluar, masuk , i float64
	fmt.Scanln(&modal)
	membeliKain(modal, &sisa, &keluar)
	totalUang -= sisa
	totalKeluar += keluar
	membeliBenangJahit(modal, &sisa, &keluar)
	totalUang -= sisa
	totalKeluar += keluar
	for i = 1; i<= 2; i++ {
		membuatPolaBaju(modal, &sisa, &keluar)
		totalUang -= sisa
		totalKeluar += keluar
	}
	for i = 1; i<=4; i++ {
		menjahitBaju(modal, &sisa, &keluar)
		totalUang -= sisa
		totalKeluar += keluar
	}
	for i = 1; i<= 2; i++ {
		mengemasBaju(modal, &sisa, &keluar)
		totalUang -= sisa
		totalKeluar += keluar
	}
	mendistribusikan(modal,&sisa,&masuk, &keluar)
	totalUang -= sisa
	totalKeluar += keluar
	fmt.Println(int(totalKeluar), int(masuk), int(totalUang)*-1)
}

func membeliKain(uangAwal float64, tUang, tPengeluaran *float64 ) {
	*tUang = uangAwal - (uangAwal/4)
	*tPengeluaran = (uangAwal/4)
}

func membeliBenangJahit(uangAwal float64, tUang, tPengeluaran *float64){
	*tUang = uangAwal - (uangAwal/20)
	*tPengeluaran = (uangAwal/20)
}

func membuatPolaBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (uangAwal/200)
	*tPengeluaran = (uangAwal/200)
}

func menjahitBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (uangAwal/200)
	*tPengeluaran = (uangAwal/200)
}

func mengemasBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - ((3*uangAwal)/200)
	*tPengeluaran = ((3*uangAwal)/200)
}

func mendistribusikan(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
	*tPengeluaran = ((3*uangAwal)/200)
	*tUang = uangAwal - ((3*uangAwal)/200)
	*tPemasukan = uangAwal/2
}