package main

import "fmt"

// 94. 二叉树的中序遍历
// 简单
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 时间复杂度：O(n)，其中 n 是二叉搜索树的节点数。每一个节点恰好被遍历一次。
// 空间复杂度：O(n)，为递归过程中栈的开销，平均情况下为 O(logn)，最坏情况下树呈现链状，为 O(n)。
func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		ret = append(ret, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	return ret
}

// 迭代
func inorderTraversal_1(root *TreeNode) []int {
	var ret []int
	var stack = []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		ret = append(ret, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return ret
}

func buildBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	var tree = &TreeNode{Val: arr[0]}
	if len(arr) == 1 {
		return tree
	}

	var index = 1
	var queue = []*TreeNode{tree}
	for len(queue) > 0 {
		var temp []*TreeNode
		for i := range queue {
			if index < len(arr) && arr[index] != -1 {
				queue[i].Left = &TreeNode{Val: arr[index]}
				temp = append(temp, queue[i].Left)
			}
			index++
			if index < len(arr) && arr[index] != -1 {
				queue[i].Right = &TreeNode{Val: arr[index]}
				temp = append(temp, queue[i].Right)
			}
			index++
		}
		queue = temp
	}
	return tree
}

func test01() {
	//          1
	//       /    \
	//   null      2
	//            /
	//           3
	fmt.Println("inorderTraversal:", inorderTraversal(buildBinaryTree([]int{1, -1, 2, 3}))) // -1代表null
	fmt.Println("inorderTraversal:", inorderTraversal(buildBinaryTree([]int{})))
	fmt.Println("inorderTraversal:", inorderTraversal(buildBinaryTree([]int{1})))

	fmt.Println("inorderTraversal_1:", inorderTraversal_1(buildBinaryTree([]int{1, -1, 2, 3}))) // -1代表null
	fmt.Println("inorderTraversal_1:", inorderTraversal_1(buildBinaryTree([]int{})))
	fmt.Println("inorderTraversal_1:", inorderTraversal_1(buildBinaryTree([]int{1})))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
