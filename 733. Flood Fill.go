package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		dfs(image, sr, sc, color, newColor)
	}
	return image
}

func dfs(image [][]int, r, c, color, newColor int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != color {
		return
	}
	image[r][c] = newColor
	dfs(image, r-1, c, color, newColor)
	dfs(image, r, c-1, color, newColor)
	dfs(image, r+1, c, color, newColor)
	dfs(image, r, c+1, color, newColor)
}
