package main

import (
	"fmt"
	"math"
)

// 7. 整数反转
// 中等

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−2^31, 2^31 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。

// 时间复杂度：O(log∣x∣)。翻转的次数即 x 十进制的位数。
// 空间复杂度：O(1)。
func reverse(x int) int {
	var ret int
	for x != 0 {
		ret = ret*10 + x%10
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}

		x /= 10
	}
	return ret
}

func test01() {
	fmt.Println("reverse:", reverse(123))
	fmt.Println("reverse:", reverse(-123))
	fmt.Println("reverse:", reverse(120))
	fmt.Println("reverse:", reverse(0))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
