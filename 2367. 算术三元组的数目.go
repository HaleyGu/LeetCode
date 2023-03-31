package main

import "fmt"

// 2367. 算术三元组的数目
// 简单

// 给你一个下标从 0 开始、严格递增 的整数数组 nums 和一个正整数 diff 。如果满足下述全部条件，则三元组 (i, j, k) 就是一个 算术三元组 ：

// i < j < k ，
// nums[j] - nums[i] == diff 且
// nums[k] - nums[j] == diff
// 返回不同 算术三元组 的数目。

// 哈希集合
// 时间复杂度 O(n)。其中 n 为数组 nums 的长度
// 空间复杂度 O(n)。
func arithmeticTriplets(nums []int, diff int) int {
	if len(nums) < 3 {
		return 0
	}
	var vis = make(map[int]bool)
	for _, v := range nums {
		vis[v] = true
	}

	var cnt int
	for _, v := range nums {
		if vis[v+diff] && vis[v+diff+diff] {
			cnt++
		}
	}
	return cnt
}

func test01() {
	fmt.Println("arithmeticTriplets:", arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	fmt.Println("arithmeticTriplets:", arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
