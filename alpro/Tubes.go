package main

import "fmt"

const NMAX int = 1000
const mkMAX int = 10
type mahasiswa struct {
	nama string
	nim int
	matkulS [mkMAX] matkul
	nMatkul int
}
type matkul struct {
	namaMatkul string
	uts, uas, quiz, total float64
	grade string
}

type listMhsw [NMAX] mahasiswa
var tabMhsw listMhsw
var nMhsw int

func main() {
	buatMenu()
}

func buatMenu() {
	/* I.S : -
	   FS: menampilkan menu utama*/
	var x int
	fmt.Println("===========================")
	fmt.Println(" APLIKASI NILAI MAHASISWA")
	fmt.Println("===========================")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Matakuliah")
	fmt.Println("3. Lihat list")
	fmt.Println("4. Cetak Transkrip Nilai")
	fmt.Println("5. Keluar")
	fmt.Println("---------------------------")	
	fmt.Print("pilihan: ") 
	fmt.Scan(&x)
	if x == 1 {
		dataMahasiswa()
	}else if x == 2 {
		dataMatakuliah()
	}else if x == 3 {

	}
}

func dataMahasiswa() {
	/*IS : -
	FS: menampilkan tampilan menu data mahasiswa*/
	var x int
	fmt.Println("===========================")
	fmt.Println("     Data MAHASISWA")
	fmt.Println("===========================")
	fmt.Println("1. Tambah Data Mahasiswa")
	fmt.Println("2. Ubah Data Mahasiswa")
	fmt.Println("3. List Data Mahasiswa")
	fmt.Println("4. Hapus Data Mahasiswa")
	fmt.Println("5. Kembali")
	fmt.Println("---------------------------")	
	fmt.Print("pilihan: ") 
	fmt.Scan(&x)
	if x == 1 {
		tambahMahasiswa(&tabMhsw)
	}else if x == 2 {
		ubahMahasiswa(&tabMhsw)
	}else if x == 3 {
		lihatList(tabMhsw)
	}else if x == 4 {
		hapusMahasiswa(&tabMhsw)
	}else if x == 5{
		buatMenu()
	}
}

func tambahMahasiswa(A *listMhsw) {
	/*IS :variabel A terefinisi sembarang 
	 FS: menambahkan data berupa Nama mahasiswa dan NIM mahasiswa berdasarkan input user*/
	var pilih string
	fmt.Println("--------------------------")
	fmt.Print("masukkan Nama: "); fmt.Scan(&A[nMhsw].nama)
	fmt.Print("masukkan NIM: "); fmt.Scan(&A[nMhsw].nim)
	nMhsw++
	fmt.Print("tambah data lagi?(y/n)"); fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi "); fmt.Scan(&pilih)
	}
	if pilih == "y" {
		tambahMahasiswa(A)
	}else if pilih == "n" {
		dataMahasiswa()
	}
}

func ubahMahasiswa(A *listMhsw) {
	/*IS : array A listMhsw terdefinisi
	 proses: mencari nama mahasiswa akan diganti
	  FS : data array A berubah sesuai inputan user */
	var targetNama, pilih string
	var targetNIM int
	var i int = 0
	var found bool = false
	fmt.Println("---------------------")
	fmt.Print("NIM mahasiswa: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found{
		if targetNIM == A[i].nim{
			fmt.Print("ubah nama: "); fmt.Scan(&targetNama)
			fmt.Print("ubah NIM: "); fmt.Scan(&targetNIM)
			A[i].nama = targetNama
			A[i].nim = targetNIM
			found = true
		}
		i++
	}
	if found == false {
		fmt.Println("data tidak ditemukan")
	}
	fmt.Print("ubah data lagi?(y/n) "); fmt.Scan(&pilih)
	if pilih == "y" {
		ubahMahasiswa(A)
	}else {
		dataMahasiswa()
	}
}

func hapusMahasiswa(A *listMhsw) {
	var pilih string
	var targetNIM int
	var i int = 0
	var found bool = false
	fmt.Println("--------------------------")
	fmt.Print("NIM Mahasiswa: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found {
		if targetNIM == A[i].nim {
			for j := i; j < nMhsw;j++ {
				A[j] = A[j+1]
			}
			found = true
			nMhsw--
		}
		i++
	}
	if found == false {
		fmt.Println("data tidak ditemukan")
	}
	fmt.Print("ubah data lagi?(y/n) "); fmt.Scan(&pilih)
	if pilih == "y" {
		hapusMahasiswa(A)
	}else {
		dataMahasiswa()
	}
}

func dataMatakuliah() {
	var x int
	fmt.Println("===========================")
	fmt.Println("     Data MATA KULIAH")
	fmt.Println("===========================")
	fmt.Println("1. Tambah Data Mata Kuliah")
	fmt.Println("2. Ubah Data Mata Kuliah")
	fmt.Println("3. List Data Mahasiswa")
	fmt.Println("4. Hapus Data Mata Kuliah")
	fmt.Println("5. Kembali")
	fmt.Println("---------------------------")	
	fmt.Print("pilihan: ") 
	fmt.Scan(&x)
	if x == 1 {
		tambahMatkul(&tabMhsw)
	}else if x == 2 {
		ubahMatkul(&tabMhsw)
	}else if x == 3{
		Listmatkul(tabMhsw)
	}else if x == 4 {
		hapusMatkul(&tabMhsw)

	}else if x == 5 {
		buatMenu()
	}
}

func tambahMatkul(A *listMhsw) {
	var i, targetNIM int
	var targetNAMA, pilih string
	var found bool
	fmt.Println("--------------------------")
	fmt.Print("masukkan nama: "); fmt.Scan(&targetNAMA)
	fmt.Print("masukkan nim: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found {
		if A[i].nama == targetNAMA && A[i].nim == targetNIM {
			fmt.Print("masukkan nama mata kuliah: ");fmt.Scan(&A[i].matkulS[A[i].nMatkul].namaMatkul)
			fmt.Print("masukkan nilai quiz: ");fmt.Scan(&A[i].matkulS[A[i].nMatkul].quiz)
			fmt.Print("masukkan nilai UTS: ");fmt.Scan(&A[i].matkulS[A[i].nMatkul].uts)
			fmt.Print("masukkan nilai UAS: ");fmt.Scan(&A[i].matkulS[A[i].nMatkul].uas)
			A[i].matkulS[A[i].nMatkul].total = addTotal(A[i].matkulS[A[i].nMatkul].uas, A[i].matkulS[A[i].nMatkul].quiz, A[i].matkulS[A[i].nMatkul].uts)
			A[i].matkulS[A[i].nMatkul].grade = addGrade(A[i].matkulS[A[i].nMatkul].total)
			A[i].nMatkul++
			found = true
		}
		i++
	}
	if found == false {
		fmt.Println("mahasiswa tidak ditemukan")
	}
	fmt.Print("tambah data lagi?(y/n) ");fmt.Scan(&pilih)
	if pilih == "y" {
		tambahMatkul(A)
	}else {
		dataMatakuliah()
	}
}

func addGrade(nilai float64) string {
	var grade string
	if nilai > 80.0 {
		grade = "A"
	}else if nilai <= 80 && nilai > 70 {
		grade = "AB"
	}else if nilai >65 && nilai <=70 {
		grade = "B"
	}else if nilai > 60 && nilai <= 65 {
		grade = "BC"
	}else if nilai > 50 && nilai <= 60 {
		grade = "C"
	}else if nilai > 40 && nilai <= 50 {
		grade = "D"
	}else {
		grade = "E"
	}
	return grade
}

func addTotal(uas, quiz, uts float64) float64{
	return (uts+uas+quiz)/3
}

func ubahMatkul(A *listMhsw) {
	var targetNAMA, targetMatkul, pilih string
	var targetNIM int
	var targetUTS, targetUAS, targetQuiz float64
	var i, j int = 0, 0
	var found1, found2 bool = false, false
	fmt.Println("---------------------")
	fmt.Print("nama mahasiswa: "); fmt.Scan(&targetNAMA)
	fmt.Print("NIM mahasiswa: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found1{
		if targetNAMA == A[i].nama && targetNIM == A[i].nim{
			fmt.Print("mata kuliah yang ingin di ubah: "); fmt.Scan(&targetMatkul)
			for j < A[i].nMatkul && !found2 {
				if targetMatkul == A[i].matkulS[j].namaMatkul {
					fmt.Print("nilai UTS yang benar: ");fmt.Scan(&targetUTS)
					fmt.Print("nilai UAS yang benar: ");fmt.Scan(&targetUAS)
					fmt.Print("nilai quiz yang benar: ");fmt.Scan(&targetQuiz)
					A[i].matkulS[j].uts = targetUTS
					A[i].matkulS[j].uas = targetUAS
					A[i].matkulS[j].quiz = targetQuiz
					A[i].matkulS[j].total = addTotal(A[i].matkulS[j].uas,A[i].matkulS[j].quiz,A[i].matkulS[j].uts)
					A[i].matkulS[j].grade = addGrade(A[i].matkulS[j].total)
					found2 = true
				}
				j++
			}
			found1 = true
		}
		i++
	}
	if found1 == false{
		fmt.Println("Mahasiswa tidak ditemukan")
	}
	if found2 == false {
		fmt.Println("mata kuliah tidak ditemukan")
	}
	fmt.Print("ubah nilai lagi?(y/n) "); fmt.Scan(&pilih)
	if pilih == "y" {
		ubahMatkul(A)
	}else {
		dataMatakuliah()
	}
}

func hapusMatkul(A *listMhsw) {
	var targetNama, pilih, targetMatkul string
	var targetNIM int
	var i,j int = 0,0
	var found1, found2 bool = false, false
	fmt.Println("--------------------------")
	fmt.Print("nama Mahasiswa: "); fmt.Scan(&targetNama)
	fmt.Print("NIM Mahasiswa: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found1 {
		if targetNama == A[i].nama && targetNIM == A[i].nim {
			fmt.Print("mata kuliah yang dihapus: ");fmt.Scan(&targetMatkul)
			for j < A[i].nMatkul && !found2 {
				if targetMatkul == A[i].matkulS[j].namaMatkul {
					for k:= j; k < A[i].nMatkul; k++ {
						A[k] = A[k+1]
					}
					found2 = true
					A[i].nMatkul--
				}
				j++
			}
			found1 = true
		}
		i++
	}
	if found1 == false {
		fmt.Println("mahasiswa tidak ditemukan")
	}
	if found2 == false {
		fmt.Println("mata kuliah tidak ditemukan")
	}
	fmt.Print("hapus mata kuliah lagi?(y/n)"); fmt.Scan(&pilih)
	if pilih == "y" {
		hapusMatkul(A)
	}else {
		dataMatakuliah()
	}
}


func lihatList(A listMhsw) {
	var i int
	var pilih string
	fmt.Println("=========================")
	fmt.Println("      list mahasiswa 	")
	fmt.Println("-------------------------")
	for i = 0; i < nMhsw; i++ {
		fmt.Println(A[i].nama, A[i].nim)
	}
	fmt.Println("=========================")
	fmt.Print("kembali?(y/n) "); fmt.Scan(&pilih)
	if pilih == "y" {
		dataMahasiswa()
	}
}

func Listmatkul(A listMhsw) {
	var i,j  int = 0, 0
	var pilih string
	var targetNama string
	var targetNIM int
	var found bool = false
	fmt.Print("Nama mahasiswa: ");fmt.Scan(&targetNama)
	fmt.Print("NIM mahasiswa: ");fmt.Scan(&targetNIM)
	for i < nMhsw && !found {
		if A[i].nama == targetNama && A[i].nim == targetNIM {
			fmt.Println("=========================")
			fmt.Println("      list matkul 	")
			fmt.Printf(" Nama: %s\n", targetNama)
			fmt.Printf(" NIM: %d \n", targetNIM)
			fmt.Println("-------------------------")
			for j < A[i].nMatkul {
				fmt.Println(A[i].matkulS[j].namaMatkul, A[i].matkulS[j].quiz,A[i].matkulS[j].uts,A[i].matkulS[j].uas, A[i].matkulS[j].grade)
				j++
			}
			fmt.Println("=========================")
			found = true
		}
		i++
	}
	fmt.Print("kembali?(y/n) "); fmt.Scan(&pilih)
	if pilih == "y" {
		dataMatakuliah()
	}
}

