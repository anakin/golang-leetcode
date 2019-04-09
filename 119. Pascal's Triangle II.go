package main

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	for k := range ret {
		ret[k] = 1
	}
	for i := 0; i < rowIndex-1; i++ {
		for j := i + 1; j >= 1; j-- {
			ret[j] = ret[j] + ret[j-1]
		}
	}
	return ret
}
