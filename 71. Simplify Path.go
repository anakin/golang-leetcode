package main

import "strings"

func simplifyPath(path string) string {
	dirs := strings.FieldsFunc(path, func(s rune) bool {
		if s == '/' {
			return true
		}
		return false
	})
	for i := 0; i < len(dirs); i++ {
		if dirs[i] == "." {
			dirs = append(dirs[:i], dirs[i+1:]...)
			i--
		} else {
			if dirs[i] == ".." {
				if i == 0 {
					dirs = dirs[1:]
					i--
					continue
				}
				dirs = append(dirs[:i-1], dirs[i+1:]...)
				i -= 2
			}
		}
	}
	res := "/"
	for i := 0; i < len(dirs); i++ {
		res += dirs[i]
		if i != len(dirs)-1 {
			res += "/"
		}
	}
	return res
}
