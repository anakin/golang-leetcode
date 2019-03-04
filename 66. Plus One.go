package main

func plusOne(digits []int) []int {

	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	//如果有进位
	res := []int{1}
	res = append(res, digits...)
	return res
}
