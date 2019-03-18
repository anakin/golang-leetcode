package main

func reverseStr(s string, k int) string {
	t := []byte(s)
	for l := 0; ; l += 2 * k {
		r := l + k - 1
		if r >= len(t) {
			r = len(t) - 1
		}
		for i := l; i < r; i, r = i+1, r-1 {
			t[i], t[r] = t[r], t[i]
		}
		if (l + 2*k) >= len(t) {
			break
		}
	}
	return string(t)
}
