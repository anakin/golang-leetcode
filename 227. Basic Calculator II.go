package main

import (
	"unicode"
)

/**
需要处理的就是乘除的情况
设置一个栈，如果是+就直接入栈，如果是-，就负数入栈
如果是乘或者除，就从栈里取出最后一个元素，运算之后入栈
最后从栈里取出所有元素相加
需要注意的是处理多位数字的情况
*/

func calculate(s string) int {
	num := 0
	st := []int{}
	var sign byte
	sign = '+'
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) { //如果是多位数字
			num = num*10 + int(s[i]-'0')
		}
		//如果当前字符是运算符，或者是最后一个字符（一定是数字）
		if s[i] != ' ' && !unicode.IsDigit(rune(s[i])) || i == len(s)-1 {
			switch sign {
			case '+':
			case '-':
				num = -num
			case '*':
				num = st[len(st)-1] * num
				st = st[:len(st)-1]
			case '/':
				num = st[len(st)-1] / num
				st = st[:len(st)-1]
			}
			sign = s[i]
			st = append(st, num)
			num = 0
		}
	}
	for _, v := range st {
		num += v
	}
	return num
}
