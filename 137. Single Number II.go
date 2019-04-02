package main

/**
这是一个简单的思路
貌似可以用位运算搞定。。。
*/
func singleNumber(nums []int) int {
	k := make(map[int]int)
	// count=0
	for _, kvs := range nums {
		if _, ok := k[kvs]; ok {
			k[kvs] += 1

		} else {
			k[kvs] = 1
		}
	}
	p := 0
	for l := range k {
		if k[l] == 1 {
			p = l
		}
	}
	return p
}
