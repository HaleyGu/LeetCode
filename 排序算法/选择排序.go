package main

import "fmt"

// 选择排序
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
// 不稳定排序
func SelectSort(arr []int) []int {
	arrLen := len(arr)
	left, right := 0, arrLen-1
	for left <= right {
		minPos, maxPos := left, left
		for i := left + 1; i <= right; i++ {
			if arr[i] < arr[minPos] {
				minPos = i
			}
			if arr[i] > arr[maxPos] {
				maxPos = i
			}
		}

		arr[left], arr[minPos] = arr[minPos], arr[left]
		if maxPos == left {
			maxPos = minPos
		}
		arr[right], arr[maxPos] = arr[maxPos], arr[right]
		left++
		right--
	}

	return arr
}

func test01() {
	fmt.Println("SelectSort:", SelectSort([]int{10, 5, 11, 9, 0}))
	fmt.Println("SelectSort:", SelectSort([]int{-1, -5, 0, 0, 456, 1}))
	fmt.Println("SelectSort:", SelectSort([]int{1, 2, 3, 4, 5}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
