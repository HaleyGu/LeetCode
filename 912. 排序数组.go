package main

import (
	"fmt"
)

// 912. 排序数组
// 中等
// https://leetcode.cn/problems/sort-an-array/

// 给你一个整数数组 nums，请你将该数组升序排列。

func merge(left, right []int) []int {
	m, n := len(left), len(right)
	i, j := 0, 0
	var ret []int
	for {
		for i < m && left[i] <= right[j] {
			ret = append(ret, left[i])
			i++
		}
		if i == m {
			ret = append(ret, right[j:]...)
			break
		}
		for j < n && right[j] < left[i] {
			ret = append(ret, right[j])
			j++
		}
		if j == n {
			ret = append(ret, left[i:]...)
			break
		}
	}
	return ret
}

func sortArrayByMerge(nums []int, start, end int) []int {
	if end-start < 2 {
		return nums[start:end]
	}
	mid := (end + start) / 2
	left := sortArrayByMerge(nums, start, mid)
	right := sortArrayByMerge(nums, mid, end)
	return merge(left, right)
}

// func sortArrayByQuick(nums []int, start, end int) []int {
// 	if start >= end {
// 		return nums
// 	}

// 	var (
// 		mid  = start + (end-start)/2
// 		i, j = start, end
// 	)

// 	for i < j {
// 		for i < j && nums[i] < nums[mid] {
// 			i++
// 		}
// 		for i < j && nums[j] >= nums[mid] {
// 			j--
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	nums[i], nums[mid] = nums[mid], nums[i]
// 	sortArrayByQuick(nums, start, i)
// 	sortArrayByQuick(nums, i+1, end)
// 	return nums
// }

func sortArrayByQuick(nums []int) []int {
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		// 递归终止条件
		if left > right {
			return nil
		}

		// 左右指针及主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}
			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}
			// 交换i/j下标元素
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换元素
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)
		quick(nums, i+1, right)
		return nums
	}

	return quick(nums, 0, len(nums)-1)
}

func sortArray(nums []int) []int {
	// return sortArrayByMerge(nums, 0, len(nums))
	return sortArrayByQuick(nums)
}

func test01() {
	fmt.Println("sortArray:", sortArray([]int{2, 2, 2, 2, 2, 2, 2, 2}))
	fmt.Println("sortArray:", sortArray([]int{5, 2, 3, 1}))
	fmt.Println("sortArray:", sortArray([]int{5, 1, 1, 2, 0, 0}))
	fmt.Println("sortArray:", sortArray([]int{-2, 3, -5}))
	fmt.Println("sortArray:", sortArray([]int{-1, 2, -8, -10}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
