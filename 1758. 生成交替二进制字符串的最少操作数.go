package main

import "fmt"

// 1758. 生成交替二进制字符串的最少操作数
// 简单

// 给你一个仅由字符 '0' 和 '1' 组成的字符串 s 。一步操作中，你可以将任一 '0' 变成 '1' ，或者将 '1' 变成 '0' 。
// 交替字符串 定义为：如果字符串中不存在相邻两个字符相等的情况，那么该字符串就是交替字符串。例如，字符串 "010" 是交替字符串，而字符串 "0100" 不是。
// 返回使 s 变成 交替字符串 所需的 最少 操作数。

// 时间复杂度：O(n)，其中 n 为输入 s 的长度，仅需遍历一遍字符串。
// 空间复杂度：O(1)，只需要常数额外空间。

func minOperations(s string) int {
	var cnt int
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != i%2 {
			cnt++
		}
	}

	if len(s)-cnt < cnt {
		return len(s) - cnt
	}
	return cnt
}

func test01() {
	fmt.Println("minOperations:", minOperations("0100"))
	fmt.Println("minOperations:", minOperations("10"))
	fmt.Println("minOperations:", minOperations("1111"))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
