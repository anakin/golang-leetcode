package main

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/
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
