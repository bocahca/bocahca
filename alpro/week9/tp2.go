package main

import "fmt"


func main() {
	const NMAX int = 52
	type TabInt [NMAX] int 
	var kartu TabInt
	var j, n int
	n = 0
	fmt.Scan(&kartu[0])
	for kartu[n] != 0 {
		n++
		fmt.Scan(&kartu[n])
	}

	for j = n-1 ; j >= 0 ; j-- {
		fmt.Print(kartu[j], " ")
	}
}
