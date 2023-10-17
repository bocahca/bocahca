package main

import "fmt"

const NMAX int = 1000
const mkMAX int = 10

type mahasiswa struct {
	nama    string
	nim     int
	matkulSiswa [mkMAX] dataMatkul
	nMatkul int
	IPK float64
	totalSKS int
	nilaiSearch float64 // variable untuk menampung total nilai 
}
type dataMatkul struct {
	nama, kode string
	sks int
	quiz, uts, uas, total float64
	grade string
}
type matkul struct {
	namaMatkul string
	kodeMatkul string
	sks int 
}

type listMhsw [NMAX]mahasiswa
type listMatkul [NMAX]matkul

var tabMhsw listMhsw
var tabMatkul listMatkul
var nMhsw, nMatkul, nNilai int

func main() {
	buatMenu()
}

func buatMenu() {
	/* I.S : -
	   FS: menampilkan menu utama*/
	var x int
	fmt.Println("================================================================================================")
	fmt.Println()
	fmt.Println(" 	 	 		APLIKASI NILAI MAHASISWA")
	fmt.Println()
	fmt.Println("================================================================================================")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Matakuliah")
	fmt.Println("3. Ambil Mata Kuliah")
	fmt.Println("4. cari")
	fmt.Println("5. Cetak Transkrip")
	fmt.Println("6. Keluar")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("pilihan: ")
	fmt.Scan(&x)
	if x != 1 && x != 2 && x !=3 && x != 4 && x != 5 && x != 6 {
		fmt.Print("pilihan tidak valid, pilih lagi: ");fmt.Scan(&x)
	}
	if x == 1 {
		dataMahasiswa()
	} else if x == 2 {
		dataMatakuliah()
	} else if x == 3 {
		ambilMatkul()
	}else if x == 4 {
		cari()
	}else if x == 5 {
		cetakTranskrip(tabMhsw)
	}
}

func dataMahasiswa() {
	/*IS : -
	FS: menampilkan tampilan menu data mahasiswa*/
	var x int
	fmt.Println("================================================================================================")
	fmt.Println("    					Data MAHASISWA")
	fmt.Println("================================================================================================")
	fmt.Println("1. Tambah Data Mahasiswa")
	fmt.Println("2. Ubah Data Mahasiswa")
	fmt.Println("3. Hapus Data Mahasiswa")
	fmt.Println("4. Kembali")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("pilihan: ")
	fmt.Scan(&x)
	if x == 1 {
		tambahMahasiswa(&tabMhsw)
	} else if x == 2 {
		ubahMahasiswa(&tabMhsw)
	} else if x == 3 {
		hapusMahasiswa(&tabMhsw)
	} else if x == 4 {
		buatMenu()
	}
}

func searchMhsw(A listMhsw, x int) int{
	/* mengembalikan nilai false jika nim mahasiswa yang dicari tidak ditemukan dan true ketika ditemukan */
	var idx int = -1
	for i:= 0; i < nMhsw; i++ {
		if A[i].nim == x {
			idx = i
		}
	}
	return idx
}
func tambahMahasiswa(A *listMhsw) {
	/*IS :variabel A terefinisi sembarang
	FS: menambahkan data berupa Nama mahasiswa dan NIM mahasiswa berdasarkan input user. Jika mahasiswa sudah terdaftar, input tidak dimasukkan ke dalam array A*/
	var pilih, nama string
	var nim int
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("masukkan Nama: ")
	fmt.Scan(&nama)
	fmt.Print("masukkan NIM: ")
	fmt.Scan(&nim)
	if searchMhsw(*A,nim) == -1 {
		A[nMhsw].nama = nama
		A[nMhsw].nim = nim
		nMhsw++
	}else {
		fmt.Println("mahasiswa sudah terdaftar")
	}	
	fmt.Print("tambah data lagi?(y/n)")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		tambahMahasiswa(A)
	} else if pilih == "n" {
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
	listmahasiswa(*A)
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("NIM mahasiswa: ")
	fmt.Scan(&targetNIM)
	for i < nMhsw && !found {
		if targetNIM == A[i].nim {
			fmt.Print("ubah nama: ")
			fmt.Scan(&targetNama)
			fmt.Print("ubah NIM: ")
			fmt.Scan(&targetNIM)
			A[i].nama = targetNama
			A[i].nim = targetNIM
			found = true
		}
		i++
	}
	if found == false {
		fmt.Println("data tidak ditemukan")
	}
	fmt.Print("ubah data lagi?(y/n) ")
	fmt.Scan(&pilih)
	if pilih == "y" {
		ubahMahasiswa(A)
	} else {
		dataMahasiswa()
	}
}

func hapusMahasiswa(A *listMhsw) {
	var pilih string
	var targetNIM int
	var i int = 0
	var found bool = false
	listmahasiswa(*A)
	fmt.Println("-------------------------------------------------------------")
	fmt.Print("NIM Mahasiswa: ")
	fmt.Scan(&targetNIM)
	for i < nMhsw && !found {
		if targetNIM == A[i].nim {
			for j := i; j < nMhsw; j++ {
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
	fmt.Print("ubah data lagi?(y/n) ")
	fmt.Scan(&pilih)
	if pilih == "y" {
		hapusMahasiswa(A)
	} else {
		dataMahasiswa()
	}
}

func dataMatakuliah() {
	var x int
	fmt.Println("================================================================================================")
	fmt.Println("    		  		 Data MATA KULIAH")
	fmt.Println("================================================================================================")
	fmt.Println("1. Tambah Data Mata Kuliah")
	fmt.Println("2. Ubah Data Mata Kuliah")
	fmt.Println("3. Hapus Data Mata Kuliah")
	fmt.Println("4. Kembali")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("pilihan: ")
	fmt.Scan(&x)
	if x != 1 && x != 2 && x !=3 && x != 4 {
		fmt.Print("pilihan tidak valid, pilih lagi: ");fmt.Scan(&x)
	}
	if x == 1 {
		tambahMatkul(&tabMatkul)
	} else if x == 2 {
		ubahMatkul(&tabMatkul)
	} else if x == 3 {
		hapusMatkul(&tabMatkul)
	} else if x == 4 {
		buatMenu()
	}
}
func searchKode(A listMatkul, x string) int {
	/* mengembalikan nilai true apabila kode matkul ditemukan */
	var idx int
	idx = -1
	for i:= 0; i < nMatkul; i++ {
		if A[i].kodeMatkul == x {
			idx = i
		}
	}
	return idx
}
func tambahMatkul(A *listMatkul){
	var pilih, nama, kode string
	var sks int
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("masukkan Nama matkul: ")
	fmt.Scan(&nama)
	fmt.Print("masukkan Kode matkul: ")
	fmt.Scan(&kode)
	fmt.Print("masukkan jumlah SKS: ")
	fmt.Scan(&sks)
	if searchKode(*A,kode) == -1 {
		A[nMatkul].kodeMatkul = kode
		A[nMatkul].namaMatkul = nama
		A[nMatkul].sks = sks
		nMatkul++
	}else {
		fmt.Println("mata kuliah sudah terdaftar")
	}		
	fmt.Print("tambah data lagi?(y/n)")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		tambahMatkul(A)
	} else if pilih == "n" {
		dataMatakuliah()
	}
}

func ubahMatkul(A *listMatkul) {
	var targetNama,targetKode, pilih string
	var targetSKS int
	var i int = 0
	var found bool = false
	ListMatkul(*A)
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("Kode Mata Kuliah: ")
	fmt.Scan(&targetKode)
	for i < nMatkul && !found {
		if targetKode == A[i].kodeMatkul {
			fmt.Print("ubah nama mata kuliah: ")
			fmt.Scan(&targetNama)
			fmt.Print("ubah SKS: ")
			fmt.Scan(&targetSKS)
			fmt.Print("ubah kode: ")
			fmt.Scan(&targetKode)
			A[i].namaMatkul = targetNama
			A[i].sks = targetSKS
			A[i].kodeMatkul = targetKode
			found = true
		}
		i++
	}
	if found == false {
		fmt.Println("data tidak ditemukan")
	}
	fmt.Print("ubah data lagi?(y/n) ")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		ubahMatkul(A)
	} else if pilih == "n" {
		dataMatakuliah()
	}
}

func hapusMatkul(A *listMatkul) {
	var pilih, targetKode string
	var i int = 0
	var found bool = false
	ListMatkul(*A)
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("Kode Matkul:  ")
	fmt.Scan(&targetKode)
	for i < nMatkul && !found {
		if targetKode == A[i].kodeMatkul {
			for j := i; j < nMhsw; j++ {
				A[j] = A[j+1]
			}
			found = true
			nMatkul--
		}
		i++
	}
	if found == false {
		fmt.Println("data tidak ditemukan")
	}
	fmt.Print("ubah data lagi?(y/n) ")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		hapusMatkul(A)
	} else if pilih == "n" {
		dataMatakuliah()
	}
}

func ListMatkul(A listMatkul) {
	var i int
	fmt.Println("================================================================================================")
	fmt.Println("    	 	  			list mata Kuliah	")
	fmt.Println("================================================================================================")
	fmt.Println("Nama					", "Kode					", "jumlah_SKS					")
	for i = 0; i < nMatkul; i++ {
		fmt.Println(A[i].namaMatkul,"					",A[i].kodeMatkul,"					", A[i].sks)
	}
}


func listmahasiswa(A listMhsw) {
	var i int
	fmt.Println("================================================================================================")
	fmt.Println("					list mahasiswa ")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("Nama						","NIM")
	for i = 0; i < nMhsw; i++ {
		fmt.Println(A[i].nama,"					", A[i].nim)
	}
}

func ambilMatkul() {
	var x int
	fmt.Println("=============================================================")
	fmt.Println("    		 Ambil Mata Kuliah ")
	fmt.Println("=============================================================")
	fmt.Println("1. Tambah matkul")
	fmt.Println("2. Ubah nilai matkul")
	fmt.Println("3. hapus matkul")
	fmt.Println("4. kembali")
	fmt.Println("-------------------------------------------------------------")
	fmt.Print("pilihan: ")
	fmt.Scan(&x)
	if x != 1 && x != 2 && x !=3 && x != 4 {
		fmt.Print("pilihan tidak valid, pilih lagi: ");fmt.Scan(&x)
	}
	if x == 1 {
		tambahNilai(&tabMhsw, tabMatkul)
	}else if x == 2 {
		ubahNilai(&tabMhsw)
	} else if x == 3 {
		hapusNilai(&tabMhsw)
	} else if x == 4 {
		buatMenu()
	} 
}

func tambahNilai(A *listMhsw, B listMatkul) {
	var targetNIM int
	var targetKode, pilih string
	var i int = 0	
	var found bool = false
	fmt.Println("------------------------------------------------------------------------------------------------")
	listmahasiswa(*A)
	fmt.Print("masukkan NIM mahasiswa: ")
	fmt.Scan(&targetNIM)
	ListMatkul(tabMatkul)
	fmt.Print("masukkan kode mata kuliah: ")
	fmt.Scan(&targetKode)
	for i < nMhsw && !found{
		if A[i].nim == targetNIM && searchKode(tabMatkul,targetKode) !=-1 {
			A[i].matkulSiswa[A[i].nMatkul].nama = B[searchKode(tabMatkul,targetKode)].namaMatkul
			A[i].matkulSiswa[A[i].nMatkul].sks = B[searchKode(tabMatkul,targetKode)].sks
			A[i].matkulSiswa[A[i].nMatkul].kode = B[searchKode(tabMatkul,targetKode)].kodeMatkul
			fmt.Print("masukkan nilai quiz: "); fmt.Scan(&A[i].matkulSiswa[A[i].nMatkul].quiz)
			fmt.Print("masukkan nilai UTS: "); fmt.Scan(&A[i].matkulSiswa[A[i].nMatkul].uts)
			fmt.Print("masukkan nilai UAS: "); fmt.Scan(&A[i].matkulSiswa[A[i].nMatkul].uas)
			A[i].matkulSiswa[A[i].nMatkul].total = addTotal(A[i].matkulSiswa[A[i].nMatkul].uas, A[i].matkulSiswa[A[i].nMatkul].quiz, A[i].matkulSiswa[A[i].nMatkul].uts)
			A[i].matkulSiswa[A[i].nMatkul].grade = addGrade(A[i].matkulSiswa[A[i].nMatkul].total)
			A[i].totalSKS = A[i].totalSKS + A[i].matkulSiswa[A[i].nMatkul].sks
			A[i].nMatkul++
			found = true
		}
		i++
	}
	if found == false {
		fmt.Println("data tidak dapat ditambahkan")
	}
	fmt.Print("tambah data lagi?(y/n) ")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		tambahNilai(A, B)
	} else if pilih == "n" {
		ambilMatkul()
	}
}

func addTotal(uas, quiz, uts float64) float64{
	return ((uts+uas+quiz)/3)
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

func ubahNilai(A *listMhsw) {
	var targetKode,pilih string
	var targetNIM int
	var targetUTS, targetUAS, targetQuiz float64
	var i, j int = 0, 0
	var found1, found2 bool = false, false
	fmt.Println("---------------------------------")
	listmahasiswa(*A)
	fmt.Print("NIM mahasiswa: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found1{
		if targetNIM == A[i].nim{
			ListMatkul(tabMatkul)
			fmt.Print("kode matkul yang ingin diubah: "); fmt.Scan(&targetKode)
			for j < A[i].nMatkul && !found2 {
				if targetKode == A[i].matkulSiswa[j].kode {
					fmt.Print("nilai UTS yang benar: ");fmt.Scan(&targetUTS)
					fmt.Print("nilai UAS yang benar: ");fmt.Scan(&targetUAS)
					fmt.Print("nilai quiz yang benar: ");fmt.Scan(&targetQuiz)
					A[i].matkulSiswa[j].uts = targetUTS
					A[i].matkulSiswa[j].uas = targetUAS
					A[i].matkulSiswa[j].quiz = targetQuiz
					A[i].matkulSiswa[j].total = addTotal(A[i].matkulSiswa[j].uas,A[i].matkulSiswa[j].quiz,A[i].matkulSiswa[j].uts)
					A[i].matkulSiswa[j].grade = addGrade(A[i].matkulSiswa[j].total)
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
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		ubahNilai(A)
	} else if pilih == "n" {
		ambilMatkul()
	}
}

func hapusNilai(A *listMhsw) {
	var pilih, targetKode string
	var targetNIM int
	var i,j int = 0,0
	var found1, found2 bool = false, false
	fmt.Println("-------------------------------------------------------------")
	listmahasiswa(*A)
	fmt.Print("NIM Mahasiswa: "); fmt.Scan(&targetNIM)
	for i < nMhsw && !found1 {
		if targetNIM == A[i].nim {
			ListMatkul(tabMatkul)
			fmt.Print("Kode mata kuliah yang dihapus: ");fmt.Scan(&targetKode)
			for j < A[i].nMatkul && !found2 {
				if targetKode == A[i].matkulSiswa[j].kode {
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
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		hapusNilai(A)
	} else if pilih == "n" {
		ambilMatkul()
	}
}

func cari() {
	var x int
	fmt.Println("================================================================================================")
	fmt.Println("    					search")
	fmt.Println("================================================================================================")
	fmt.Println("1. Cari mata kuliah")
	fmt.Println("2. Cari mahasiswa")
	fmt.Println("3. daftar mahasiswa")
	fmt.Println("4. Kembali")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Print("pilihan: ")
	fmt.Scan(&x)
	if x == 1 {
		cariMatkul(tabMhsw)
	} else if x == 2 {
		cariMahasiswa(tabMhsw)
	} else if x == 3 {
		daftarMahasiswa(tabMhsw)
	} else if x == 4 {
		buatMenu()
	}
}
func sorttotalSKS(A *listMhsw) {
	/*IS : array A terdefinisi
	  FS : array a terurut secara descending berdasarkan total sks yang diambil */
	var i, pass int
	var key mahasiswa
	for pass = 1; pass < nMhsw; pass++ {
		key = A[pass]
		i = pass - 1
		for i >= 0 && A[i].totalSKS < key.totalSKS {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}
func daftarMahasiswa(A listMhsw) {
	/*IS : Array A terdefinisi
	FS : daftar mahasiswa tercetak secara descending berdasarkan SKS yang di ambil */
	var pilih string
	fmt.Println("================================================================================================")
  	fmt.Println("    					Daftar Mahasiswa")
	fmt.Println("================================================================================================")
	sorttotalSKS(&A)
	fmt.Println("NAMA		","NIM		","TOTAL_SKS")
	for i := 0; i < nMhsw; i++ {
		fmt.Println(A[i].nama,"		",A[i].nim,"		",A[i].totalSKS)
	}
	fmt.Print("kembali ke menu?"); fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		cari()
	}
}

func cariMahasiswa(A listMhsw) {
/*IS : array A terdefinisi
  FS: menampilkan nama dan NIM mahasiswa yang dicari, serta mata kuliah yang diambil */
	var targetNIM, idx int
	var pilih string
	fmt.Print("Masukkan NIM mahasiswa: ") ;fmt.Scan(&targetNIM)
	idx = searchMhsw(A, targetNIM)
	if idx != -1 {
		fmt.Println("================================================================================================")
		fmt.Printf("    					Nama: %s\n", A[idx].nama)
		fmt.Printf("    					NIM: %d\n", A[idx].nim)
		fmt.Println("------------------------------------------------------------------------------------------------")
		fmt.Println("Nama Matkul","		","Kode Matkul","		","Jumlah SKS")
		fmt.Println()
		for i:=0; i< A[idx].nMatkul; i++ {
			fmt.Println(A[idx].matkulSiswa[i].nama, "			",A[idx].matkulSiswa[i].kode,"			",A[idx].matkulSiswa[i].sks)
		}
		fmt.Println("------------------------------------------------------------------------------------------------")
		fmt.Println("Total SKS: ",A[idx].totalSKS)
		fmt.Println("------------------------------------------------------------------------------------------------")
	}
	fmt.Print("Kembali ke menu?(y/n) ")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		cari()
	}
}
func sortNilai(A *listMhsw) {
/*IS: array A terdefinisi 
  FS: array A terurut secara descending berdasarkan total nilai */
	var i, pass int
	var key mahasiswa
	for pass = 1; pass < nNilai; pass++ {
		key = A[pass]
		i = pass - 1
		for i >= 0 && A[i].nilaiSearch < key.nilaiSearch {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}
func cariMatkul(A listMhsw) {
/* IS: array mahasiswa A terdefinisi
   FS: menampilkan daftar mahasiswa yang mengambil Mata kuliah tertentu secara descending berdasarkan total nilai */
	var newArr listMhsw
	var targetKode, pilih string 
	var i, j, idx int
	nNilai = 0
	fmt.Print("masukkan kode mata kuliah: "); fmt.Scan(&targetKode)
	idx = searchKode(tabMatkul, targetKode)
	for i = 0; i < nMhsw; i++ {
		for j = 0; j < A[i].nMatkul; j++ {
			if A[i].matkulSiswa[j].kode == targetKode {
				newArr[nNilai] = A[i]
				newArr[nNilai].nilaiSearch = A[i].matkulSiswa[j].total
				nNilai++
			}
		}
	}
	if idx != -1 {
		sortNilai(&newArr)
		fmt.Println("================================================================================================")
		fmt.Printf("Nama Mata Kuliah: %s\n", tabMatkul[idx].namaMatkul)
		fmt.Printf("Kode: %s\n", tabMatkul[idx].kodeMatkul)
		fmt.Println("------------------------------------------------------------------------------------------------")
		fmt.Println("Nama			","NIM			","Total Nilai")
		for i=0; i< nNilai; i++ {
			fmt.Printf("%s			%d			%.2f\n",newArr[i].nama,	newArr[i].nim, newArr[i].nilaiSearch)
		}
		fmt.Println("------------------------------------------------------------------------------------------------")
	}else {
		fmt.Print("Mata kuliah tidak ditemukan")
	}
	fmt.Print("Kembali ke menu?(y/n) ")
	fmt.Scan(&pilih)
	for pilih != "y" && pilih != "n" {
		fmt.Print("input tidak valid, silahkan pilih lagi ")
		fmt.Scan(&pilih)
	}
	if pilih == "y" {
		cari()
	}
}

func countIPK(A listMhsw, x int) float64 {
	var i, idx, j int
	var IPMatkul, IPK float64
	for i = 0; i < nMhsw; i++ {
		if A[i].nim == x {
			idx = i
		}
	}
	for j = 0; j < A[idx].nMatkul; j++ {
		if A[idx].matkulSiswa[j].grade == "A" {
			IPMatkul = 4.00 * float64(A[idx].matkulSiswa[j].sks)
		}else if A[idx].matkulSiswa[j].grade == "AB" {
			IPMatkul = 3.50 * float64(A[idx].matkulSiswa[j].sks)
		}else if A[idx].matkulSiswa[j].grade == "B" {
			IPMatkul = 3.00 * float64(A[idx].matkulSiswa[j].sks)
		}else if A[idx].matkulSiswa[j].grade == "BC" {
			IPMatkul = 2.50 * float64(A[idx].matkulSiswa[j].sks)
		}else if A[idx].matkulSiswa[j].grade == "C" {
			IPMatkul = 2.00 * float64(A[idx].matkulSiswa[j].sks)
		}else if A[idx].matkulSiswa[j].grade == "D" {
			IPMatkul = 1.00 * float64(A[idx].matkulSiswa[j].sks)
		}else if A[idx].matkulSiswa[j].grade == "E" {
			IPMatkul = 0 * float64(A[idx].matkulSiswa[j].sks)
		}
		IPK = IPK + IPMatkul
	}
	IPK = IPK / float64(A[idx].totalSKS)
	return IPK
}
func cetakTranskrip(A listMhsw) {
	var targetNIM int
	var pilih string
	fmt.Print("Masukkan NIM mahasiswa: "); fmt.Scan(&targetNIM)
	for i := 0; i < nMhsw; i++ {
		if A[i].nim == targetNIM {
			A[i].IPK = countIPK(A, targetNIM)
			fmt.Println("===================================================================================================================")
			fmt.Println("    					TRANSKRIP NILAI")
			fmt.Printf("Nama: %s\n", A[i].nama)
			fmt.Printf("NIM: %d\n", A[i].nim)
			fmt.Println("-------------------------------------------------------------------------------------------------------------------")
			fmt.Println("no	","Kode matkul		","Nama Matkul			","jml SKS		","grade		")
			for j:= 0; j < A[i].nMatkul; j++ {
				fmt.Printf("%d	 %s			 %s				 %d			 %s			\n", j+1, A[i].matkulSiswa[j].kode, A[i].matkulSiswa[j].nama, A[i].matkulSiswa[j].sks, A[i].matkulSiswa[j].grade)
			}
			fmt.Println("-------------------------------------------------------------------------------------------------------------------")
			fmt.Printf("IPK Mahasiswa: %.2f\n", A[i].IPK)
			fmt.Println("-------------------------------------------------------------------------------------------------------------------")
		}else {
			fmt.Println("mahasiswa tidak ditemukan")
		}
		fmt.Print("Kembali ke menu?(y/n) ")
		fmt.Scan(&pilih)
		for pilih != "y" && pilih != "n" {
			fmt.Print("input tidak valid, silahkan pilih lagi ")
			fmt.Scan(&pilih)
		}
		if pilih == "y" {
			buatMenu()
		}
	}
}