package main

import "fmt"

// 200. 岛屿数量
// 中等

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

// 时间：O(mn)
// 空间：O(1)
func numIslands(grid [][]byte) int {
	var m, n = len(grid), len(grid[0])
	// 方向：前 后 左 右
	var direction = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var landNum int

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, dire := range direction {
			if i+dire[0] >= 0 && i+dire[0] < m && j+dire[1] < n && j+dire[1] >= 0 {
				dfs(i+dire[0], j+dire[1])
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				landNum++
				dfs(i, j)
			}
		}
	}
	return landNum
}

func test01() {
	fmt.Println("numIslands:", numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println("numIslands:", numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
