package main

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	ret := ""
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	for num > 0 {
		a := num % 7
		ret = strconv.Itoa(a) + ret
		num = num / 7
	}
	if flag {
		ret = "-" + ret
	}
	return ret
}
