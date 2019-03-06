package main

func countSegments(s string) int {

	count := 1
	found := false
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == ' ' && s[i+1] != ' ' {
			count++
		}
		if s[i] != ' ' {
			found = true
		}
		if i == len(s)-1 && s[i] == ' ' {
			count--
		}
	}
	if found {
		return count
	}
	return 0
}
