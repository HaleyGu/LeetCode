package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 468. 验证IP地址
// 中等
// https://leetcode.cn/problems/validate-ip-address/

// 给定一个字符串 queryIP。如果是有效的 IPv4 地址，返回 "IPv4" ；如果是有效的 IPv6 地址，返回 "IPv6" ；如果不是上述类型的 IP 地址，返回 "Neither" 。

// 有效的IPv4地址 是 “x1.x2.x3.x4” 形式的IP地址。 其中 0 <= xi <= 255 且 xi 不能包含 前导零。
// 例如: “192.168.1.1” 、 “192.168.1.0” 为有效IPv4地址， “192.168.01.1” 为无效IPv4地址; “192.168.1.00” 、 “192.168@1.1” 为无效IPv4地址。

// 一个有效的IPv6地址 是一个格式为“x1:x2:x3:x4:x5:x6:x7:x8” 的IP地址，其中: 1 <= xi.length <= 4
// xi 是一个 十六进制字符串 ，可以包含数字、小写英文字母( 'a' 到 'f' )和大写英文字母( 'A' 到 'F' )。
// 在 xi 中允许前导零。
// 例如 "2001:0db8:85a3:0000:0000:8a2e:0370:7334" 和 "2001:db8:85a3:0:0:8A2E:0370:7334" 是有效的 IPv6 地址，
// 而 "2001:0db8:85a3::8A2E:037j:7334" 和 "02001:0db8:85a3:0000:0000:8a2e:0370:7334" 是无效的 IPv6 地址。

// 时间复杂度：O(n)，其中 n 是字符串 queryIP 的长度。我们只需要遍历字符串常数次。
// 空间复杂度：O(1)。
func validIPAddress(queryIP string) string {
	var Neither = "Neither"
	var IPv4 = "IPv4"
	var IPv6 = "IPv6"

	if ipv4 := strings.Split(queryIP, "."); len(ipv4) == 4 {
		for _, v := range ipv4 {
			if (len(v) > 1 && v[0] == '0') || len(v) > 3 {
				return Neither
			}
			if i, err := strconv.Atoi(v); err != nil || i < 0 || i > 255 {
				return Neither
			}
		}
		return IPv4
	}
	if ipv6 := strings.Split(queryIP, ":"); len(ipv6) == 8 {
		for _, v := range ipv6 {
			if len(v) > 4 {
				return Neither
			}
			if _, err := strconv.ParseInt(v, 16, 64); err != nil {
				return Neither
			}
		}
		return IPv6
	}
	return Neither
}

func test01() {
	fmt.Println("validIPAddress:", validIPAddress("172.16.254.1"))
	fmt.Println("validIPAddress:", validIPAddress("256.256.256.256"))
	fmt.Println("validIPAddress:", validIPAddress("192.168.1.00"))
	fmt.Println("validIPAddress:", validIPAddress("192.168@.1.1"))
	fmt.Println("validIPAddress:", validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println("validIPAddress:", validIPAddress("2001:0db8:85a3::8A2E:037j:7334"))
	fmt.Println("validIPAddress:", validIPAddress("2001:db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println("validIPAddress:", validIPAddress("02001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
