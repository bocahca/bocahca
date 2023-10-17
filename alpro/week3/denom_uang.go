package main

import "fmt"

func main() {
	var uang, l10, l2, l1 int
	fmt.Scanln(&uang)
	pecahUang(uang, &l10, &l2, &l1)
	fmt.Println(l10, "lembar 10000")
	fmt.Println(l2, "lembar 2000")
	fmt.Println(l1, "lembar 1000")
}
func pecahUang(money int, k10, k2, k1 *int) {
	var temp int
	*k10 = money / 10000
	temp = money % 10000
	*k2 = temp / 2000
	temp = temp % 2000
	*k1 = temp / 1000
}
