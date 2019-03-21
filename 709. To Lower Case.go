package main

func toLowerCase(str string) string {
	s := []byte(str)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] += 'a' - 'A'
		}
	}
	return string(s)
}
