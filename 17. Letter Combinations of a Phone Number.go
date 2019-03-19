package main

/*
回溯算法
*/
func letterCombinations(digits string) []string {
	table := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ret := []string{}
	if len(digits) > 0 {
		help(&ret, digits, "", 0, table)
	}
	return ret
}

func help(ret *[]string, digits string, cur string, index int, table []string) {
	if index == len(digits) {
		*ret = append(*ret, cur)
		return
	}
	tmp := table[digits[index]-48]
	for _, t := range tmp {
		help(ret, digits, cur+string(t), index+1, table)
	}
}
