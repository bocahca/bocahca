package main

import "fmt"

const NMAX int = 1000

type himpunan struct {
	info [NMAX] string
	nElemen int
}

func main() {
	var A, B, C, D himpunan
	createSet_1303220115(&A)
	createSet_1303220115(&B)
	intersection_1303220115(A, B, &C)
	union_1303220115(A,B, &D)
	printSet_1303220115(C)
	fmt.Println()
	printSet_1303220115(D)
}

func createSet_1303220115(set *himpunan) {
	var j int
	var sama bool
	fmt.Scan(&set.info[0])
	set.nElemen = 1
	sama = false
	for set.nElemen < NMAX && !sama {
		fmt.Scan(&set.info[set.nElemen])
		for j = 0; j < set.nElemen; j++ {
			if set.info[set.nElemen] == set.info[j] {
				sama = true
				set.nElemen--
			}
		}
		set.nElemen++
	}
} 

func printSet_1303220115(set himpunan) {
	var i int
	for i = 0; i < set.nElemen; i++ {
		fmt.Print(set.info[i], " ")
	}
}

func isMember_1303220115(set himpunan, s string) bool {
	var i int
	var ketemu bool
	ketemu = false
	for i < set.nElemen && !ketemu {
		ketemu = set.info[i] == s
		i++
	}
	return ketemu
}

func intersection_1303220115(set1, set2 himpunan, set3 *himpunan) {
	var i,k int
	for i < set1.nElemen {
		if isMember_1303220115(set2,set1.info[i]) {
			set3.info[k] = set1.info[i]
			set3.nElemen++
			k++
		}
		i++
	}
}

func union_1303220115(set1, set2 himpunan, set3 *himpunan) {
	var k int
	for set3.nElemen < set1.nElemen {
		set3.info[set3.nElemen] = set1.info[set3.nElemen]
		set3.nElemen++
	}
	for k = 0; k <=set2.nElemen + 1; k++ {
		if !isMember_1303220115(set1,set2.info[k]) {
			set3.nElemen++
			set3.info[set3.nElemen] = set2.info[k]

		}
	} 
}
