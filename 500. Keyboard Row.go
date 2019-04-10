package main

/**
先把每一行用map保存，然后查找每个词
*/
func findWords(words []string) []string {
	str := []string{"QWERTYUIOPqwertyuiop", "ASDFGHJKLasdfghjkl", "ZXCVBNMzxcvbnm"}
	dic := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			dic[str[i][j]] = i
		}
	}
	ret := make([]string, 0, len(words))
	for _, word := range words {
		flag := true
		for i := 0; flag && i < len(word)-1; i++ {
			if dic[word[i]] != dic[word[i+1]] {
				flag = false
			}
		}
		if flag {
			ret = append(ret, word)
		}
	}
	return ret
}
