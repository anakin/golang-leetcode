package main

func reverseVowels(s string) string {
	t := []byte(s)
	m := make(map[byte]bool)
	m['o'], m['e'], m['a'], m['i'], m['u'] = true, true, true, true, true
	m['O'], m['E'], m['A'], m['I'], m['U'] = true, true, true, true, true
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		for i < j && !m[t[i]] {
			i++
		}
		for j > i && !m[t[j]] {
			j--
		}
		if i >= j {
			break
		}
		t[i], t[j] = t[j], t[i]
	}
	return string(t)

}
