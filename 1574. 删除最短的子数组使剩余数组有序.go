package main

import "fmt"

// 1574. 删除最短的子数组使剩余数组有序
// 中等

// 给你一个整数数组 arr ，请你删除一个子数组（可以为空），使得 arr 中剩下的元素是 非递减 的。
// 一个子数组指的是原数组中连续的一个子序列。
// 请你返回满足题目要求的最短子数组的长度。

// 时间复杂度：O(n)，其中 n 是 arr 的长度。双指针 i 和 j 移动过程中，每个元素最多只会被遍历两次，所以复杂度为 O(n)。
// 空间复杂度：O(1)。
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	i, j := 0, n-1

	for i < n-1 && arr[i] <= arr[i+1] {
		i++
	}
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if i >= j {
		return 0
	}

	var ret = min(j, n-i-1)

	for m := 0; m <= i; m++ {
		for j < n && arr[m] > arr[j] {
			j++
		}
		ret = min(ret, j-m-1)
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func test01() {
	fmt.Println("findLengthOfShortestSubarray:", findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println("findLengthOfShortestSubarray:", findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println("findLengthOfShortestSubarray:", findLengthOfShortestSubarray([]int{1, 2, 3}))
	fmt.Println("findLengthOfShortestSubarray:", findLengthOfShortestSubarray([]int{1}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
