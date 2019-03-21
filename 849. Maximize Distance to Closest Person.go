package main

/**
如果两头是0，开头和末尾就是距离
如果中间是0，距离是0的个数/2
取最大值
*/
func maxDistToClosest(seats []int) int {
	l := len(seats)
	ret := []int{}
	for k, v := range seats {
		if v == 1 {
			ret = append(ret, k)
		}
	}
	m := len(ret)
	pre := ret[0]
	max := mymax(pre, l-1-ret[m-1])
	for _, v := range ret {
		max = mymax(max, (v-pre)/2)
		pre = v
	}
	return max
}

func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
