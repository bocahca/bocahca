package main

import "fmt"

func main() {
	var tabInt [100] int
	var n, i int
	fmt.Scanln(&n) 
	for i = 0; i < n; i++ {
		fmt.Scanln(&tabInt[i])
	}
	cetakArr(tabInt, n)
	insSort(&tabInt,n)
	cetakArr(tabInt,n)
}

func insSort(A *[100]int, N int) {
	var pass, i, key int
	for pass = 1; pass < N; pass++ {
		key = A[pass]
		i = pass - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}

func cetakArr(A [100]int, n int) {
	for i := 0; i< n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}
