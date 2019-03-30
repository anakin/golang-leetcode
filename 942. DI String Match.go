package main

/**
发现D的时候，就放最大的值进去，发现I的时候，就放最小的值进去
*/

func diStringMatch(S string) []int {
	l := len(S)
	ret := make([]int, l+1)
	indexI, indexD := 0, len(S)
	for i := 0; i <= l; i++ {
		if i == l {
			ret[i] = indexI
			break
		}
		if S[i] == 'D' {
			ret[i] = indexD
			indexD--
		}
		if S[i] == 'I' {
			ret[i] = indexI
			indexI++
		}
	}
	return ret
}
