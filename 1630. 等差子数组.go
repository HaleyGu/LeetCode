package main

import (
	"fmt"
	"sort"
)

// 1630. 等差子数组
// 中等

// 如果一个数列由至少两个元素组成，且每两个连续元素之间的差值都相同，那么这个序列就是 等差数列 。
// 更正式地，数列 s 是等差数列，只需要满足：对于每个有效的 i ， s[i+1] - s[i] == s[1] - s[0] 都成立。
// 例如，下面这些都是 等差数列 ：
// 1, 3, 5, 7, 9
// 7, 7, 7, 7
// 3, -1, -5, -9

// 下面的数列 不是等差数列 ：
// 1, 1, 2, 5, 7

// 给你一个由 n 个整数组成的数组 nums，和两个由 m 个整数组成的数组 l 和 r，后两个数组表示 m 组范围查询，其中第 i 个查询对应范围 [l[i], r[i]] 。
// 所有数组的下标都是 从 0 开始 的。
// 返回 boolean 元素构成的答案列表 answer 。
// 如果子数组 nums[l[i]], nums[l[i]+1], ... , nums[r[i]] 可以 重新排列 形成 等差数列 ，answer[i] 的值就是 true；否则answer[i] 的值就是 false 。

// 时间O(m*n)
// 空间O(n)
func check(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	var temp = make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)

	for i := 2; i < len(nums); i++ {
		if temp[i]-temp[i-1] == temp[1]-temp[0] {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var ret = make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		ret[i] = check(nums[l[i] : r[i]+1])
	}
	return ret
}

func test01() {
	fmt.Println("checkArithmeticSubarrays:", checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println("checkArithmeticSubarrays:", checkArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
