package main

/**
思路就是每一个为1的点，看它的上下左右是否也为1，若为1则有一边重复，最后加的时候要减去重复的边。
*/

func islandPerimeter(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				t := 4
				if j >= 1 && grid[i][j-1] == 1 { //左边有邻居
					t--
				}
				if j < col-1 && grid[i][j+1] == 1 { //右边有邻居
					t--
				}
				if i >= 1 && grid[i-1][j] == 1 { //上面有邻居
					t--
				}
				if i < row-1 && grid[i+1][j] == 1 { //下面有邻居
					t--
				}
				count += t
			}
		}
	}
	return count
}
