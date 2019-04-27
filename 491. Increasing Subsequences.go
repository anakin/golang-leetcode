package main

var res [][]int

func isUsed(nums []int, start, index int) bool {
	for i := start; i < index; i++ {
		if nums[i] == nums[index] {
			return true
		}
	}
	return false
}

func findSubsequences(nums []int) [][]int {
	res = [][]int{}
	dfs(nums, 0, []int{})
	return res

}

func dfs(nums []int, start int, tmp []int) {
	if len(tmp) > 1 {
		res = append(res, tmp)
	}
	if start == len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		if len(tmp) > 0 && nums[i] < tmp[len(tmp)-1] {
			continue
		}
		if i > start && isUsed(nums, start, i) {
			continue
		}
		tmp = append(tmp, nums[i])
		dfs(nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
