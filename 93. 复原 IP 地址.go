package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 93. 复原 IP 地址
// 中等
// https://leetcode.cn/problems/restore-ip-addresses/

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
// 你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

// 时间复杂度：O(3^SEG_COUNT ×∣s∣)。由于 IP 地址的每一段的位数不会超过 3，因此在递归的每一层，我们最多只会深入到下一层的 3 种情况。
// 由于 SEG_COUNT=4，对应着递归的最大层数，所以递归本身的时间复杂度为 O(3^SEG_COUNT)。
// 如果我们复原出了一种满足题目要求的 IP 地址，那么需要 O(∣s∣) 的时间将其加入答案数组中，因此总时间复杂度为 O(3^SEG_COUNT ×∣s∣)。

// 空间复杂度：O(SEG_COUNT)，这里只计入除了用来存储答案数组以外的额外空间复杂度。
// 递归使用的空间与递归的最大深度 SEG_COUNT 成正比。并且在上面的代码中，我们只额外使用了长度为 SEG_COUNT 的数组 segments 存储已经搜索过的 IP 地址，因此空间复杂度为 O(SEG_COUNT)。
func restoreIpAddresses(s string) []string {
	var SEG_COUNT = 4
	var sLen = len(s)
	if sLen < SEG_COUNT {
		return nil
	}

	var ret []string
	var ip []string

	var dfs func(s string, start int)
	dfs = func(s string, start int) {
		if len(ip) == SEG_COUNT && start == sLen {
			ret = append(ret, strings.Join(ip, "."))
			return
		} else if len(ip) == SEG_COUNT && start != sLen {
			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > sLen {
				return
			}
			temp := s[start : start+i]
			if i != 1 && temp[0] == '0' {
				continue
			}
			if v, _ := strconv.Atoi(temp); v > 255 || v < 0 {
				continue
			}

			ip = append(ip, temp)
			dfs(s, start+i)
			ip = ip[:len(ip)-1]
		}
	}

	dfs(s, 0)
	return ret
}

func test01() {
	fmt.Println("restoreIpAddresses:", restoreIpAddresses("25525511135"))
	fmt.Println("restoreIpAddresses:", restoreIpAddresses("0000"))
	fmt.Println("restoreIpAddresses:", restoreIpAddresses("101023"))
	fmt.Println("restoreIpAddresses:", restoreIpAddresses("000"))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
