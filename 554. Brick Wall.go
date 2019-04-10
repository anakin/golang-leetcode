package main

/**
我们可以拿砖缝左侧砖块的总长度来标记每个砖缝，这样遍历每一行的砖块将所有砖缝位置计数后存入Hash Table中，最终遍历Hash Table找出同一纵向位置砖缝最多的地方即可。墙的总行数减同一纵向位置的最大砖缝数即是最少需要凿穿的砖块数
*/

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	maxend := 0
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j+1 < len(wall[i]); j++ {
			sum += wall[i][j]
			if v, ok := m[sum]; ok {
				m[sum] = v + 1
			} else {
				m[sum] = 1
			}
			maxend = max(maxend, m[sum])
		}
	}
	return len(wall) - maxend
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
