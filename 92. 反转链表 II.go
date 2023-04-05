package main

import "fmt"

// 92. 反转链表 II
// 中等

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 一次遍历「穿针引线」反转链表（头插法）
// 时间复杂度：O(N)，其中 N 是链表总节点数。最多只遍历了链表一次，就完成了反转。
// 空间复杂度：O(1)。只使用到常数个变量。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{Next: head}
	var leftNode = dummy
	for i := 0; i < left-1; i++ {
		leftNode = leftNode.Next
	}

	var curr = leftNode.Next
	for i := 0; i < right-left; i++ {
		temp := curr.Next
		curr.Next = temp.Next
		temp.Next, leftNode.Next = leftNode.Next, temp
	}

	return dummy.Next
}

func buildList(arr []int) *ListNode {
	var dummy = &ListNode{}
	var curr = dummy
	for _, v := range arr {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func visitList(list *ListNode) []int {
	var arr []int
	var curr = list
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}

func test01() {
	fmt.Println("reverseBetween:", visitList(reverseBetween(buildList([]int{1, 2, 3, 4, 5}), 2, 4)))
	fmt.Println("reverseBetween:", visitList(reverseBetween(buildList([]int{5}), 1, 1)))
	fmt.Println("reverseBetween:", visitList(reverseBetween(buildList([]int{3, 5}), 1, 2)))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
