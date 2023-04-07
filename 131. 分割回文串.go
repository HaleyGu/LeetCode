package main

import (
	"fmt"
)

// 131. 分割回文串
// 中等
// https://leetcode.cn/problems/palindrome-partitioning/

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

// 回文串 是正着读和反着读都一样的字符串。

// 回溯
// 时间复杂度：O(n*2^n)，其中 n 是字符串 s 的长度。在最坏情况下，s 包含 n 个完全相同的字符，因此它的任意一种划分方法都满足要求。
// 而长度为 n 的字符串的划分方案数为 2^(n−1)=O(2n)，每一种划分方法需要 O(n) 的时间求出对应的划分结果并放入答案，因此总时间复杂度为 O(n*2^n)。

// 空间复杂度：O(n2)，这里不计算返回答案占用的空间。
// 数组 f 需要使用的空间为 O(n2)，而在回溯的过程中，我们需要使用 O(n) 的栈空间以及 O(n) 的用来存储当前字符串分割方法的空间。
// 由于 O(n) 在渐进意义下小于 O(n2)，因此空间复杂度为 O(n2)。
func partition(s string) [][]string {
	var sLen = len(s)

	var ret [][]string
	var path []string

	var dfs func(index int)
	dfs = func(index int) {
		if index == sLen {
			ret = append(ret, append([]string{}, path...))
			return
		}

		for i := index; i < sLen; i++ {
			if !isPalindrome(s, index, i) {
				continue
			}

			path = append(path, s[index:i+1])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ret
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func test01() {
	fmt.Println("partition:", partition("aab"))
	fmt.Println("partition:", partition("a"))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
