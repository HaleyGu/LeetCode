package main

import "fmt"

// 3. 无重复字符的最长子串
// 中等

// 滑动窗口
// 时间复杂度：O(N)，其中 N 是字符串的长度。左指针和右指针分别会遍历整个字符串一次。
// 空间复杂度：O(∣Σ∣)，其中 Σ 表示字符集（即字符串中可以出现的字符），∣Σ∣ 表示字符集的大小。
// 在本题中没有明确说明字符集，因此可以默认为所有 ASCII 码在 [0,128) 内的字符，即 ∣Σ∣=128。
// 我们需要用到哈希集合来存储出现过的字符，而字符最多有 ∣Σ∣ 个，因此空间复杂度为 O(∣Σ∣)。
func lengthOfLongestSubstring(s string) int {
	var win = make(map[byte]int)
	var ret int
	var left, right = -1, 0

	for right < len(s) {
		win[s[right]]++
		for win[s[right]] > 1 {
			left++
			win[s[left]]--
		}
		ret = max(ret, right-left)
		right++
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test01() {
	fmt.Println("lengthOfLongestSubstring:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("lengthOfLongestSubstring:", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("lengthOfLongestSubstring:", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
