package main

import "fmt"

// 300. 最长递增子序列
// 中等

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

// 时间复杂度：O(n2)，其中 n 为数组 nums 的长度。
// 动态规划的状态数为 n，计算状态 dp[i] 时，需要 O(n) 的时间遍历 dp[0…i−1] 的所有状态，所以总时间复杂度为 O(n2)。
// 空间复杂度：O(n)，需要额外使用长度为 dp 数组。

// 动态规划
func lengthOfLIS(nums []int) (ans int) {
	var dp = make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	for _, d := range dp {
		ans = max(ans, d)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test01() {
	fmt.Println("lengthOfLIS:", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println("lengthOfLIS:", lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println("lengthOfLIS:", lengthOfLIS([]int{7, 7, 7, 7, 7, 7}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
