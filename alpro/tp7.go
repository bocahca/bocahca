package main

import "fmt"


func main() {
	var x, i int
	var hasil float64
	fmt.Scan(&x)
	if x == 0 {
		fmt.Println(x)
	}else {
		average_1303220115(x, i, &hasil)
		fmt.Println(hasil)
	}

}

func average_1303220115(bil, i int, hasil *float64) {
	var digit int
	if bil > 0 {
		digit = bil % 10
		*hasil = *hasil + float64(digit)
		i++
		average_1303220115(bil/10, i, hasil)
	}else  if bil == 0 {
		*hasil = *hasil / float64(i)
	}
}

