package main

/**
遍历一遍，先判断左边是不是符合规则
再判断右边
注意，如果有相等的元素，那一定是在两个部分的中间
*/
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	l := len(A)
	for i+1 < l && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == l-1 {
		return false
	}
	for i+1 < l && A[i] > A[i+1] {
		i++
	}
	return i == l-1
}

/**
由于上面的速度太慢，于是找到了最快的AC如下：
先找到最大的（就是中间的元素）

*/
func validMountainArray1(a []int) bool {
	if len(a) < 3 {
		return false
	}
	max := 0
	index := 0
	for i, v := range a {
		if v > max {
			max = v
			index = i
		}
	}
	if index == 0 || index == len(a)-1 {
		return false
	}
	for i, j := 0, 1; i <= index && j <= index; i, j = i+1, j+1 {
		if a[i] >= a[j] {
			return false
		}
	}

	for i, j := index, index+1; i < len(a) && j < len(a); i, j = i+1, j+1 {
		if a[i] <= a[j] {
			return false
		}
	}
	return true

}
