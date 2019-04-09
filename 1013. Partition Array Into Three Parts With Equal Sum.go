package main

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}
	total := 0
	for _, v := range A {
		total += v
	}
	if total%3 != 0 {
		return false
	}
	sum := 0
	for _, v := range A {
		sum += v
		if sum == total/3 {
			sum = 0
		}
	}
	return sum == 0
}
