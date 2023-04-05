package main

import (
	"fmt"
)

// 128. 最长连续序列
// 中等

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 哈希表
// 时间复杂度：O(n)，其中 n 为数组的长度。
// 空间复杂度：O(n)。哈希表存储数组中所有的数需要 O(n) 的空间。
func longestConsecutive(nums []int) int {
	var mark = make(map[int]bool)
	for _, num := range nums {
		mark[num] = true
	}

	var ret = 1
	for _, num := range nums {
		if !mark[num-1] {
			var curr = num
			var count int
			for mark[curr] {
				count++
				curr++
			}
			if ret < count {
				ret = count
			}

		}
	}
	return ret
}

func test01() {
	fmt.Println("longestConsecutive:", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println("longestConsecutive:", longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
