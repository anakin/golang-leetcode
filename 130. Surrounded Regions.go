package main

/**
TODO not AC yet
*/
func solve(board [][]byte) {
	l1 := len(board)
	if l1 == 0 {
		return
	}
	l2 := len(board[0])
	if l2 == 0 {
		return
	}
	//初始化访问数组
	visit := make([][]bool, l1)
	for i := 0; i < l1; i++ {
		visit[i] = make([]bool, l2)
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if !visit[i][j] {
				helper(board, visit, i, j)
			}
		}
	}
}

func helper(board [][]byte, visit [][]bool, x, y int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	if visit[x][y] {
		return board[x][y] != 'O'
	}
	visit[x][y] = true
	flag := true
	if board[x][y] == 'O' {
		board[x][y] = 'X'
		flag = helper(board, visit, x+1, y) && helper(board, visit, x-1, y) && helper(board, visit, x, y+1) && helper(board, visit, x, y-1)
		if flag == false {
			board[x][y] = 'O'
		}
	}
	return flag
}
