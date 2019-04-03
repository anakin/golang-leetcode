package main

/**
dfs深度优先搜索
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				helper(grid, i, j)
				ret++
			}
		}
	}
	return ret
}

func helper(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		helper(grid, i-1, j)
		helper(grid, i, j-1)
		helper(grid, i+1, j)
		helper(grid, i, j+1)
	}
}
