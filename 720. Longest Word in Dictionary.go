package main

import "sort"

/**
先排序
*/
func longestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]bool)
	ret := ""
	for _, w := range words {
		if len(w) == 1 || m[w[0:len(w)-1]] == true {
			if len(ret) < len(w) {
				ret = w
			}
			m[w] = true
		}
	}
	return ret
}
