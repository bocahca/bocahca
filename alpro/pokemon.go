package main

import "fmt"

type pokemon struct {
	nama                          string
	CP, HP, IV_att, IV_def, IV_HP float64
}

type arrPokemon [2023]pokemon

func main() {
	var T arrPokemon
	var n int
	var flag, target string
	fmt.Scan(&n)
	inputPokemon_130322015(&T, n)
	printPokemon_1303220115(T, n)
	fmt.Scanln(&target, &flag)
	mengurutkanPokemon_1303220115(&T, n, flag)
	printPokemon_1303220115(T, n)
	fmt.Println(totalIV_1303220115(T, n, target))
}

func inputPokemon_130322015(T *arrPokemon, n int) {
	var i int
	for i = 0; i <= n; i++ {
		fmt.Scanln(&T[i].nama, &T[i].CP, &T[i].HP, &T[i].IV_att, &T[i].IV_def, &T[i].IV_HP)
	}
}

func printPokemon_1303220115(T arrPokemon, n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(T[i].nama, T[i].CP, T[i].HP, T[i].IV_att, T[i].IV_def, T[i].IV_HP)
	}
}

func mengurutkanPokemon_1303220115(T *arrPokemon, n int, flag string) {
	var pass, i int
	var temp pokemon
	pass = 1
	for pass < n-1 {
		if flag == "CP" {
			i = pass
			temp = T[pass]
			for i >= 0 && temp.CP < T[i-1].CP {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		} else if flag == "HP" {
			i = pass
			temp = T[pass]
			for i >= 0 && temp.HP < T[i].HP {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		} else if flag == "IV_Atk" {
			i = pass
			temp = T[pass]
			for i >= 0 && temp.IV_att < T[i-1].IV_att {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		} else if flag == "IV_Def" {
			i = pass
			temp = T[pass]
			for i >= 0 && temp.IV_def < T[i-1].IV_def {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		} else if flag == "IV_HP" {
			i = pass
			temp = T[pass]
			for i >= 0 && temp.IV_HP < T[i-1].IV_HP {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		}
	}
}

func totalIV_1303220115(T arrPokemon, n int, namaX string) float64 {
	var total float64
	total = -9999
	for i := 0; i < n; i++ {
		if T[i].nama == namaX {
			total = ((T[i].IV_att+T[i].IV_def+T[i].IV_HP) *100) /45.0
		}
	}
	return total
}
