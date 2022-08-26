package mathSkills

import (
	"math"
	"sort"
)

func Average(t []int) float64 {
	nb := 0.0
	for _, each := range t {
		nb += float64(each)
	}
	return nb / float64(len(t))
}

func AverageFloat(t []float64) float64 {
	nb := 0.0
	for _, each := range t {
		nb += float64(each)
	}
	return nb / float64(len(t))
}

func Median(t []int) float64 {
	sort.Ints(t)
	l := len(t)
	m := int(l / 2)
	if l%2 != 0 {
		return float64(t[m])
	} else {
		return (float64(t[m]) + float64(t[m-1])) / 2
	}
}

func Variance(t []int) float64 {
	newTable := []float64{}
	m := Average(t)
	for _, each := range t {
		newTable = append(newTable, (float64(each)-m)*(float64(each)-m))
	}
	return AverageFloat(newTable)
}

func Deviation(t []int) float64 {
	return math.Sqrt(Variance(t))
}
