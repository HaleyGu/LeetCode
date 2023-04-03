package main

import (
	"fmt"
)

// 11. 盛最多水的容器
// 中等
// https://leetcode.cn/problems/container-with-most-water/

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。

// 双指针
// 时间复杂度：O(N)，双指针总计最多遍历整个数组一次。
// 空间复杂度：O(1)，只需要额外的常数级别的空间。
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		if height[left] <= height[right] {
			area = max(area, height[left]*(right-left))
			left++
		} else {
			area = max(area, height[right]*(right-left))
			right--
		}
	}
	return area
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func test01() {
	fmt.Println("maxArea:", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("maxArea:", maxArea([]int{1, 1}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
