package main

import "strings"

/**
按照规则转换之后，直接和原来的单词比较
*/
func detectCapitalUse(word string) bool {

	l := strings.ToLower(word)
	u := strings.ToUpper(word)
	t := strings.ToUpper(string(word[:1])) + strings.ToLower(word[1:])
	return l == word || u == word || t == word

}
