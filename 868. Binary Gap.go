package main

func binaryGap(N int) int {
	var d, max int

	if N&(N-1) == 0 {
		return 0
	}

	for ; N > 0; N >>= 1 {
		if N&1 == 1 {
			if max == 0 {
				max = 1
			}
			if d > max {
				max = d
			}
			d = 0
		}
		if max > 0 {
			d++
		}
	}
	return max
}
