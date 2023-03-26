package main

import "fmt"

// 2395. 和相等的子数组
// 中等

// 给你一个下标从 0 开始的整数数组 nums ，判断是否存在 两个 长度为 2 的子数组且它们的 和 相等。注意，这两个子数组起始位置的下标必须 不相同 。
// 如果这样的子数组存在，请返回 true，否则返回 false 。
// 子数组 是一个数组中一段连续非空的元素组成的序列。

// 时间复杂度：O(n2)，其中 n 为数组 nums 的长度。
// 空间复杂度：O(1)
func findSubarrays(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[i+1] == nums[j]+nums[j+1] {
				return true
			}
		}
	}
	return false
}

// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。
// 空间复杂度：O(n)。即为哈希表需要使用的空间。
func findSubarraysByHash(nums []int) bool {
	sum := make(map[int]bool)
	for i := 0; i < len(nums)-1; i++ {
		if sum[nums[i]+nums[i+1]] {
			return true
		} else {
			sum[nums[i]+nums[i+1]] = true
		}
	}
	return false
}

func test01() {
	fmt.Println("findSubarrays:", findSubarrays([]int{4, 2, 4}))
	fmt.Println("findSubarrays:", findSubarrays([]int{1, 2, 3, 4, 5}))
	fmt.Println("findSubarrays:", findSubarrays([]int{0, 0, 0, 0}))
	fmt.Println("findSubarrays:", findSubarrays([]int{0, 0}))

	fmt.Println("findSubarraysByHash:", findSubarraysByHash([]int{4, 2, 4}))
	fmt.Println("findSubarraysByHash:", findSubarraysByHash([]int{1, 2, 3, 4, 5}))
	fmt.Println("findSubarraysByHash:", findSubarraysByHash([]int{0, 0, 0, 0}))
	fmt.Println("findSubarraysByHash:", findSubarraysByHash([]int{0, 0}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
