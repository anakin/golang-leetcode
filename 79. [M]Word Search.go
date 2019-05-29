package main

/**
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/

func help(b [][]byte, word string, i, j int) bool {
	if word == "" {
		return true
	}
	if b[i][j] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	b[i][j] = '#'
	if i > 0 && help(b, word[1:], i-1, j) {
		return true
	}
	if j > 0 && help(b, word[1:], i, j-1) {
		return true
	}
	if i < len(b)-1 && help(b, word[1:], i+1, j) {
		return true
	}
	if j < len(b[0])-1 && help(b, word[1:], i, j+1) {
		return true
	}
	b[i][j] = word[0]
	return false
}
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return word == ""
	}
	if word == "" {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] != board[i][j] {
				continue
			}
			if help(board, word, i, j) {
				return true
			}
		}
	}
	return false
}
