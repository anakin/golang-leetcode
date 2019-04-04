package main

/**
假如有一个数组为[a,b,c,d],那么在经过这样一轮运算后得到的结果就是[bcd,acd,abd,abc],所以我们可以构造这样两个数组[1,a,ab,abc]和[bcd,cd,d,1].这样两个数组对应项相乘就可以得到结果
*/
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	for k := range ret {
		ret[k] = 1
	}
	tmp := 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = ret[i] * tmp
		tmp *= nums[i]
	}
	return ret
}
