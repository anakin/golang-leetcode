package main

/**
二分查找
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	for low <= high {
		mid := (low + high) / 2
		if letters[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low >= len(letters) {
		return letters[0]
	} else {
		return letters[low]
	}
}
