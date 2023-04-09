package main

import (
	"fmt"
)

// 099. 最小路径之和
// 中等
// https://leetcode.cn/problems/0i0mDW/

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：一个机器人每次只能向下或者向右移动一步。

func minPathSum1(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	var ret int = -1
	var dir = [][]int{{1, 0}, {0, 1}}
	var dfs func(i, j, sum int)
	dfs = func(i, j, sum int) {
		sum = sum + grid[i][j]
		if i == m-1 && j == n-1 {
			if sum < ret || ret == -1 {
				ret = sum
			}
		}

		for _, v := range dir {
			if i+v[0] < m && j+v[1] < n {
				dfs(i+v[0], j+v[1], sum)
			}
		}
	}

	dfs(0, 0, 0)
	return ret
}

// 动态规划
// 时间复杂度：O(mn)，其中 m 和 n 分别是网格的行数和列数。需要对整个网格遍历一次，计算 dp 的每个元素的值。
// 空间复杂度：O(mn)，其中 m 和 n 分别是网格的行数和列数。创建一个二维数组 dp，和网格大小相同。
// 空间复杂度可以优化，例如每次只存储上一行的 dp 值，则可以将空间复杂度优化到 O(n)。
func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func test01() {
	fmt.Println("minPathSum:", minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println("minPathSum:", minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
