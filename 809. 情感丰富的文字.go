package main

import "fmt"

// 809. 情感丰富的文字
// 中等

// 有时候人们会用重复写一些字母来表示额外的感受，比如 "hello" -> "heeellooo", "hi" -> "hiii"。
// 我们将相邻字母都相同的一串字符定义为相同字母组，例如："h", "eee", "ll", "ooo"。
// 对于一个给定的字符串 S ，如果另一个单词能够通过将一些字母组扩张从而使其和 S 相同，我们将这个单词定义为可扩张的（stretchy）。
// 扩张操作定义如下：选择一个字母组（包含字母 c ），然后往其中添加相同的字母 c 使其长度达到 3 或以上。
// 例如，以 "hello" 为例，我们可以对字母组 "o" 扩张得到 "hellooo"，但是无法以同样的方法得到 "helloo" 因为字母组 "oo" 长度小于 3。
// 此外，我们可以进行另一种扩张 "ll" -> "lllll" 以获得 "helllllooo"。
// 如果 s = "helllllooo"，那么查询词 "hello" 是可扩张的，因为可以对它执行这两种扩张操作使得 query = "hello" -> "hellooo" -> "helllllooo" = s。
// 输入一组查询单词，输出其中可扩张的单词数量。

// 时间复杂度：O(n∣s∣+∑i∣wordsi|)，其中 n 是数组 words 的长度，∣s∣ 和 wordsi 分别是字符串 s 和数组 words 中第 i 个字符串的长度。
// 空间复杂度：O(1)。
func expend(s, word string) bool {
	m, n := len(s), len(word)
	i, j := 0, 0

	for i < m && j < n {
		if s[i] != word[j] {
			return false
		}
		ci := s[i]
		var cnti, cntj int
		for i < m && s[i] == ci {
			cnti++
			i++
		}
		for j < n && word[j] == ci {
			cntj++
			j++
		}
		if cnti < cntj || cnti > cntj && cnti < 3 {
			return false
		}
	}

	return i == m && j == n
}

func expressiveWords(s string, words []string) int {
	var cnt int
	for _, word := range words {
		fmt.Println(word)
		if expend(s, word) {
			fmt.Println(true)
			cnt++
		}
	}
	return cnt
}

func test01() {
	fmt.Println("expressiveWords:", expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
