package main

import "fmt"

// 插入排序
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
// 稳定排序
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
	return arr
}

func test01() {
	fmt.Println("InsertionSort:", InsertionSort([]int{10, 5, 11, 9, 0}))
	fmt.Println("InsertionSort:", InsertionSort([]int{-1, -5, 0, 0, 456, 1}))
	fmt.Println("InsertionSort:", InsertionSort([]int{1, 2, 3, 4, 5}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
