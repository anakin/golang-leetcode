package main

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	m := make(map[string]int)
	m1 := make(map[string]bool)
	for _, v := range banned {
		m1[v] = true
	}
	p := strings.ToLower(paragraph)
	re, _ := regexp.Compile("[,.!?;'']")
	p = re.ReplaceAllString(p, " ")
	r := strings.Split(p, " ")
	for _, v := range r {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			if _, ok := m1[v]; !ok && strings.Trim(v, " ") != "" {
				m[v] = 1
			}
		}
	}
	max := 0
	ret := " "
	for v := range m {
		if m[v] > max {
			max = m[v]
			ret = v
		}
	}
	return ret
}
