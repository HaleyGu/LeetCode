package main

import (
	"fmt"
	"strings"
)

// 6. N 字形变换
// 中等

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	var ret = make([]string, numRows)
	cycle := 2*numRows - 2
	for i, v := range s {
		pos := i % cycle
		ret[min(pos, cycle-pos)] += string(v)
	}
	return strings.Join(ret, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func test01() {
	fmt.Println("convert:", convert("PAYPALISHIRING", 3))
	fmt.Println("convert:", convert("PAYPALISHIRING", 4))
	fmt.Println("convert:", convert("A", 1))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
