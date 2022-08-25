package mathSkills

import "sort"

func Average(t []int) int {
	nb := 0
	for _, each := range t {
		nb += each
	}
	return nb / len(t)
}

func Median(t []int) int {
	sort.Ints(t)
	l := len(t)
	m := l / 2
	if l%2 != 0 {
		return t[m]
	} else {
		return (t[m] + t[m-1]) / 2
	}
}

func Variance(t []int) int {
	newTable := []int{}
	m := Average(t)
	for _, each := range t {
		newTable = append(newTable, (each-m)*(each-m))
	}
	return Average(newTable)
}

func Deviation(t []int) int {
	v := Variance(t)
	nb := 0
	for nb*nb < v {
		nb++
	}
	return nb
}
