package main

/**
生成magic string的同时 ，计算1的数量
*/

func magicalString(n int) int {
	if n < 1 {
		return 0
	}
	if n < 4 {
		return 1
	}
	magic := make([]int, n)
	magic[0] = 1
	magic[1] = 2
	magic[2] = 2
	ret, next := 1, 1
	countIndex := 2
	count := magic[countIndex]
	i := 3
	for i < n {
		for count > 0 && i < n {
			if next == 1 {
				ret++
			}
			magic[i] = next
			count--
			i++
		}
		if next == 1 {
			next = 2
		} else {
			next = 1
		}
		countIndex++
		count = magic[countIndex]
	}
	return ret
}
