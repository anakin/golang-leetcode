package main

var ret []string

func generateParenthesis(n int) []string {
	ret = []string{}
	helper("", n, n)
	return ret
}

func helper(item string, left, right int) {
	if left < 0 || right < 0 || left > right {
		return
	}
	if left == 0 && right == 0 {
		ret = append(ret, item)
		return
	}
	helper(item+"(", left-1, right)
	helper(item+")", left, right-1)

}
