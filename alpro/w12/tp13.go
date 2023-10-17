package main

import "fmt"

type student struct {
	name                            string
	math, indo, eng, sains, average float64
}

type arrStudent [2048]student

func main() {
	var data arrStudent
	var n int
	entryStudent_1303220115(&data, &n)
	calculateAverage_1303220115(&data, n)
	rangking_1303220115(&data, n)
	printTop3_1303220115(data, n)
	printMax_1303220115(data, n)
}

func entryStudent_1303220115(T *arrStudent, n *int) {
	var i int
	fmt.Scanln(&T[0].name, &T[0].math, &T[0].indo, &T[0].eng, &T[0].sains)
	*n = 1
	for T[i].name != "STOP" {
		i++
		fmt.Scanln(&T[i].name, &T[i].math, &T[i].indo, &T[i].eng, &T[i].sains)
		*n++
	}
}

func calculateAverage_1303220115(T *arrStudent, n int) {
	for i := 0; i < n; i++ {
		T[i].average = ((T[i].math + T[i].indo + T[i].eng + T[i].sains) /4)
	}
}

func max_1303220115(T arrStudent, n int, flag string) int {
	var i, idxmax int
	idxmax = 0
	if flag == "math" {
		for i = 1; i < n; i++ {
			if T[i].math > T[idxmax].math {
				idxmax = i
			}
		}
	} else if flag == "ind" {
		for i = 1; i < n; i++ {
			if T[i].indo > T[idxmax].indo {
				idxmax = i
			}
		}
	} else if flag == "eng" {
		for i = 1; i < n; i++ {
			if T[i].eng > T[idxmax].eng {
				idxmax = i
			}
		}
	} else if flag == "sains" {
		for i = 1; i < n; i++ {
			if T[i].sains > T[idxmax].sains {
				idxmax = i
			}
		}
	}
	return idxmax
}

func rangking_1303220115(T *arrStudent, n int) {
	var pass, idx, i int
	var temp student

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if T[i].average > T[idx].average {
				idx = i
			}
			i++
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
		pass++
	}
}

func printTop3_1303220115(T arrStudent, n int) {
	fmt.Printf("Rangking 1: %s rata-rata %f \n", T[0].name, T[0].average)
	fmt.Printf("Rangking 2: %s rata-rata %f \n", T[1].name, T[1].average)
	fmt.Printf("Rangking 3: %s rata-rata %f \n", T[2].name, T[2].average)
}

func printMax_1303220115(T arrStudent, n int) {
	var idx int
	idx = max_1303220115(T, n, "math")
	fmt.Printf("Nilai matematika tertinggi diraih oleh %s \n", T[idx].name)
	idx = max_1303220115(T, n, "ind")
	fmt.Printf("Nilai Bahasa Indonesia tertinggi diraih oleh %s \n", T[idx].name)
	idx = max_1303220115(T, n, "eng")
	fmt.Printf("Nilai Bahasa Inggris tertinggi diraih oleh %s \n", T[idx].name)
	idx = max_1303220115(T, n, "sains")
	fmt.Printf("Nilai pelajaran IPA/IPS tertinggi diraih oleh %s \n", T[idx].name)
}
