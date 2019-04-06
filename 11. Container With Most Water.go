package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
