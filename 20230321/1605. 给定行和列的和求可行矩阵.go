package main

import "fmt"

// 1605. 给定行和列的和求可行矩阵
// 中等

// 给你两个非负整数数组 rowSum 和 colSum ，其中 rowSum[i] 是二维矩阵中第 i 行元素的和， colSum[j] 是第 j 列元素的和。
// 换言之你不知道矩阵里的每个元素，但是你知道每一行和每一列的和。
// 请找到大小为 rowSum.length x colSum.length 的任意 非负整数 矩阵，且该矩阵满足 rowSum 和 colSum 的要求。
// 请你返回任意一个满足题目要求的二维矩阵，题目保证存在 至少一个 可行矩阵。

// 时间复杂度：O(n×m)，其中 n 和 m 分别为数组 rowSum 和 colSum 的长度，
// 主要为构造 matrix 结果矩阵的时间开销，填充 matrix 的时间复杂度为 O(n+m)。
// 空间复杂度：O(1)，仅使用常量空间。注意返回的结果数组不计入空间开销。
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	var mat = make([][]int, len(rowSum))
	for i := range rowSum {
		mat[i] = make([]int, len(colSum))
		for j := range colSum {
			mat[i][j] = min(rowSum[i], colSum[j])
			colSum[j] -= mat[i][j]
			if rowSum[i] == mat[i][j] {
				break
			}
			rowSum[i] -= mat[i][j]
		}
	}
	return mat
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func test01() {
	fmt.Println("restoreMatrix:", restoreMatrix([]int{5, 7, 10}, []int{8, 6, 8}))
	fmt.Println("restoreMatrix:", restoreMatrix([]int{8, 6, 8}, []int{1}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
