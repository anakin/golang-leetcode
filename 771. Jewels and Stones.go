package main

func numJewelsInStones(J string, S string) int {
	jstr := []byte(J)
	sstr := []byte(S)
	m := make(map[byte]bool)
	for _, v := range jstr {
		m[v] = true
	}
	num := 0
	for _, v := range sstr {
		if _, ok := m[v]; ok {
			num++
		}
	}
	return num
}
