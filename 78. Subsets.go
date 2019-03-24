package main

/**
combination:
dfs方法，注意中间结果一定要copy到新的变量中，不然会得到空值
*/
var ans [][]int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	cur := []int{}
	for i := 0; i <= len(nums); i++ {
		dfs(nums, i, 0, cur)
	}
	return ans
}
func dfs(nums []int, n, s int, cur []int) {
	if len(cur) == n {
		tmp := make([]int, n)
		copy(tmp, cur)
		ans = append(ans, tmp)
		return
	}
	for i := s; i < len(nums); i++ {
		cur = append(cur, nums[i])
		dfs(nums, n, i+1, cur)
		cur = cur[:len(cur)-1]
	}
}
