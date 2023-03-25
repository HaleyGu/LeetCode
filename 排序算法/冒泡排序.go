package main

import "fmt"

// 冒泡排序
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
// 稳定排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		// 定义 is_changed ,记录每轮发生过交换；
		var is_changed bool
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j+1] < arr[j] {
				is_changed = true

				arr[j] = arr[j+1] ^ arr[j]
				arr[j+1] = arr[j+1] ^ arr[j]
				arr[j] = arr[j+1] ^ arr[j]
			}
		}

		// 如果一整轮没交换过，则已经有序；退出排序；
		if is_changed == false {
			break
		}
	}
	return arr
}

func test01() {
	fmt.Println("BubbleSort:", BubbleSort([]int{10, 5, 11, 9, 0}))
	fmt.Println("BubbleSort:", BubbleSort([]int{-1, -5, 0, 0, 456, 1}))
	fmt.Println("BubbleSort:", BubbleSort([]int{1, 2, 3, 4, 5}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
