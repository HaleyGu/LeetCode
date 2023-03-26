package main

import "fmt"

// 归并排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
// 稳定排序
func MergeSort(arr []int, start, end int) []int {
	if end-start < 2 {
		return arr[start:end]
	}
	mid := (end + start) / 2
	left := MergeSort(arr, start, mid)
	right := MergeSort(arr, mid, end)
	return merge(left, right)
}

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

func test01() {
	fmt.Println("MergeSort:", MergeSort([]int{10, 5, 11, 9, 0}, 0, 5))
	fmt.Println("MergeSort:", MergeSort([]int{-1, -5, 0, 0, 456, 1}, 0, 6))
	fmt.Println("MergeSort:", MergeSort([]int{1, 2, 3, 4, 5}, 0, 5))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
