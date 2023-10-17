package main

import "fmt"

type senjata struct {
	jenis string
	harga, damage, ammo int
}

const SenMax int = 100

type tabSenjata [SenMax]senjata

func main() {
	var sjt tabSenjata
	var nSjt int
	var target string
	bacadata(&sjt, &nSjt)
	cetakdata(sjt, nSjt)
	findAmmoMax(sjt, nSjt, &target)
	fmt.Println(target)
	urutDescendingbyDamage(&sjt, nSjt)
	cetakdata(sjt, nSjt)

}

func bacadata(A *tabSenjata, n *int) {
	var i int
	fmt.Scanln(&*n)
	for i = 0; i <*n; i++{
		fmt.Scanln(&A[i].jenis, &A[i].harga, &A[i].damage, &A[i].ammo)
	}
}

func cetakdata(A tabSenjata, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(A[i].jenis, A[i].harga, A[i].damage, A[i].ammo)
	}
}

func findAmmoMax(A tabSenjata, n int, jenis *string ) {
	var i, idxmax int
	idxmax = 0
	i = 1
	for i < n {
		if A[i].ammo > A[idxmax].ammo {
			idxmax = i
		}
		i++
	}
	*jenis = A[idxmax].jenis
}

func urutDescendingbyDamage( A *tabSenjata, n int) {
	var idx, i, pass int
	var temp senjata
	pass = 1
	for pass < n {
		idx = pass - 1 
		i = pass
		for i < n {
			if A[i].damage > A[idx].damage {
				idx = i
			}
			i++
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
		pass++
	}
}

