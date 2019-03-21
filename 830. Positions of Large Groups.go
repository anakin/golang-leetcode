package main

func largeGroupPositions(S string) [][]int {
	ret := [][]int{}
	i, j := 0, 0
	for i < len(S) {
		for j < len(S) && S[i] == S[j] {
			j++
		}
		if j-i >= 3 {
			tmp := []int{i, j - 1}
			ret = append(ret, tmp)
		}
		i = j
	}
	return ret
}
