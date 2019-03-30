package main

/**
关键是要理解题目。。。
*/
func findLUSlength(a string, b string) int {

	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
