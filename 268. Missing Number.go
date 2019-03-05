package main

//NB
func missingNumber(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	l := len(nums)
	sum := l * (l + 1) / 2
	return sum - total
}
