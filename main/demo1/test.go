package main

//	dfs
//
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿
// 只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
// 输入：grid = [
// ['1','1','1','1','0'],
// ['1','1','0','1','0'],
// ['1','1','0','0','0'],
// ['0','0','0','0','0']
// ]
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, visited, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfs(grid, visited, i+1, j)
	dfs(grid, visited, i-1, j)
	dfs(grid, visited, i, j+1)
	dfs(grid, visited, i, j-1)
}
