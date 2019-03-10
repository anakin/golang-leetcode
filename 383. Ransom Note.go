package main

//一个入hash表，查另一个
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		c, ok := m[v]
		if ok == false || c == 0 {
			return false
		}
		m[v]--
	}
	return true
}
