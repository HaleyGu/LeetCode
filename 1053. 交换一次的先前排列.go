package main

import "fmt"

// 1053. 交换一次的先前排列
// 中等
// https://leetcode.cn/problems/previous-permutation-with-one-swap/

// 给你一个正整数数组 arr（可能存在重复的元素），请你返回可在 一次交换（交换两数字 arr[i] 和 arr[j] 的位置）后得到的、按字典序排列小于 arr 的最大排列。
// 如果无法这么操作，就请返回原数组。

// 贪心
// 时间复杂度：O(n)，其中 n 是数组 arr 的长度。查找 i 需要 O(n) 的时间复杂度，查找 j 需要 O(n) 的时间复杂度。
// 空间复杂度：O(1)。返回值不计入空间复杂度。
func prevPermOpt1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			for j := len(arr) - 1; j > i-1; j-- {
				if arr[j] < arr[i] && arr[j] != arr[j-1] {
					arr[i], arr[j] = arr[j], arr[i]
					return arr
				}
			}
		}
	}
	return arr
}

func test01() {
	fmt.Println("prevPermOpt1:", prevPermOpt1([]int{3, 2, 1}))
	fmt.Println("prevPermOpt1:", prevPermOpt1([]int{1, 1, 5}))
	fmt.Println("prevPermOpt1:", prevPermOpt1([]int{1, 9, 4, 6, 7}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
