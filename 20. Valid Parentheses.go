package main

/**
用一个slice模拟stack
遍历字符串，遇到左括号入栈，遇到右括号出战并判断
最后判断栈是否为空
*/

func isValid(s string) bool {
	str := []byte(s)
	stack := []byte{}
	for _, v := range str {
		if '(' == v || '{' == v || '[' == v {
			stack = append(stack, v)
		}
		if len(stack) == 0 {
			return false
		}
		if ')' == v {
			if stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if '}' == v {
			if stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if ']' == v {
			if stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
