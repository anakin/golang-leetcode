package main

//从两端向中间收敛
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
