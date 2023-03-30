package main

import "fmt"

// 5. 最长回文子串
// 中等

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

// 中心扩展算法
// 时间复杂度：O(n2)，其中 n 是字符串的长度。
// 长度为 1 和 2 的回文中心分别有 n 和 n−1 个，每个回文中心最多会向外扩展 O(n) 次。
// 空间复杂度：O(1)。
func longestPalindrome(s string) string {
	var sLen = len(s)
	var index int
	var palindrome string

	for index < sLen {
		left, right := index, index
		for left >= 0 && s[left] == s[index] {
			left--
		}
		for right < sLen && s[right] == s[index] {
			right++
		}
		index = right
		for left >= 0 && right < sLen && s[left] == s[right] {
			left--
			right++
		}
		if len(palindrome) < right-left-1 {
			palindrome = s[left+1 : right]
		}
	}
	return palindrome
}

func test01() {
	fmt.Println("longestPalindrome:", longestPalindrome("babad"))
	fmt.Println("longestPalindrome:", longestPalindrome("cbbd"))
	fmt.Println("longestPalindrome:", longestPalindrome("aaaaaa"))
	fmt.Println("longestPalindrome:", longestPalindrome("asbbbbsa"))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
