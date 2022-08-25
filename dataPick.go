package mathSkills

import (
	"strconv"
	"strings"
)

func DataToInt(t []byte) []int {
	table := strings.Fields(string(t))
	newTable := []int{}
	for _, each := range table {
		x, _ := strconv.Atoi(each)
		newTable = append(newTable, x)
	}
	return newTable
}
