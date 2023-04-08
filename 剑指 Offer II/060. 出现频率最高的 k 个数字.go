package main

import (
	"fmt"
	"sort"
)

// 060. 出现频率最高的 k 个数字
// 中等
// https://leetcode.cn/problems/g5c51o/

// 给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。

// 哈希+快速排序
// 时间复杂度：O(N2)，其中 N 为数组的长度。设处理长度为 N 的数组的时间复杂度为 f(N)。
// 由于处理的过程包括一次遍历和一次子分支的递归，最好情况下，有 f(N)=O(N)+f(N/2)，根据 主定理，能够得到 f(N)=O(N)。
// 最坏情况下，每次取的中枢数组的元素都位于数组的两端，时间复杂度退化为 O(N2)。
// 但由于我们在每次递归的开始会先随机选取中枢元素，故出现最坏情况的概率很低。平均情况下，时间复杂度为 O(N)。
// 空间复杂度：O(N)。哈希表的大小为 O(N)，用于排序的数组的大小也为 O(N)，快速排序的空间复杂度最好情况为 O(logN)，最坏情况为 O(N)。
func topKFrequent(nums []int, k int) []int {
	var frequent = make(map[int]int)
	for i := range nums {
		frequent[nums[i]]++
	}

	var times Times
	for num, time := range frequent {
		times = append(times, []int{num, time})
	}
	sort.Sort(times)

	var ret = make([]int, k)
	var index int
	for index < k {
		if index >= len(times) {
			break
		}
		ret[index] = times[index][0]
		index++
	}
	return ret
}

type Times [][]int

func (t Times) Len() int {
	return len(t)
}
func (t Times) Less(i, j int) bool {
	return t[i][1] > t[j][1]
}
func (t Times) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func test01() {
	fmt.Println("topKFrequent:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println("topKFrequent:", topKFrequent([]int{1}, 1))
	fmt.Println("topKFrequent:", topKFrequent([]int{1, 5, 4, 22, 5, 4, 5, 22}, 10))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
