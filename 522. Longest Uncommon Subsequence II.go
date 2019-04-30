package main

//TODO TLE
func findLUSlength(strs []string) int {
	res, n := -1, len(strs)
	for i := 0; i < n; i++ {
		if len(strs[i]) < res {
			continue
		}
		j := -1
		for j < n {
			j++
			if i != j && isSub(strs[i], strs[j]) {
				break
			}
		}
		if j == n {
			res = max(res, len(strs[i]))
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isSub(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		}
	}
	return i == len(a)
}
