package main

/**
对于给定的数组 a1,a2,a3,a4,a5,…
假设a4 < a3. 为了实现数组的单调非减，我们必须改变a4和a3其中的一个值，与此同时，为了后续计算的需要，我们应当尽可能使a4相对较小。此时，究竟是改变a3,还是a4取决于a2值的大小。我们假设a4 < a3， 那么a3 >= a2。 因此，如果a4 < a2, 那么我们将改变a4, 最好的选择就是令 a4 = a3; 否则的话，我们令a3 = a4， 但是这种情况我们并不需要考虑
*/
func checkPossibility(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	found := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if found {
				return false
			} else {
				if i-2 >= 0 && nums[i] < nums[i-2] {
					nums[i] = nums[i-1]
				}
				found = true
			}
		}
	}
	return true
}
