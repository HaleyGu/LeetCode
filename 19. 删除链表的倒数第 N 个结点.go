package main

import "fmt"

// 19. 删除链表的倒数第 N 个结点
// 中等

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

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

// 栈
// 时间复杂度：O(L)，其中 L 是链表的长度。
// 空间复杂度：O(L)，其中 L 是链表的长度。主要为栈的开销。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var arr []*ListNode
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		arr = append(arr, node)
	}

	pre := arr[len(arr)-n-1]
	if pre.Next != nil {
		pre.Next = pre.Next.Next
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
	fmt.Println("removeNthFromEnd:", visitList(removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 2)))
	fmt.Println("removeNthFromEnd:", visitList(removeNthFromEnd(buildList([]int{1}), 1)))
	fmt.Println("removeNthFromEnd:", visitList(removeNthFromEnd(buildList([]int{1, 2}), 1)))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
