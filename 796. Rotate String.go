package main

/**
找到A中和B第一个字母相同的位置，然后重新拼接成一个新的字符串，判断是否相等
*/
func rotateString(A string, B string) bool {
	if len(A) == 0 && len(B) == 0 {
		return true
	}
	if len(A) == 0 || len(B) == 0 {
		return false
	}
	a, b := []byte(A), []byte(B)
	for k, v := range a {
		if v == b[0] {
			c := a[k:]
			c = append(c, a[:k]...)
			if string(c) == string(b) {
				return true
			}
		}
	}
	return false
}
