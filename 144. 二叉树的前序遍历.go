package main

import "fmt"

// 144. 二叉树的前序遍历
// 简单
// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

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
func preorderTraversal(root *TreeNode) []int {
	var ret []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return ret
}

// 迭代
func preorderTraversal_1(root *TreeNode) []int {
	var ret []int
	var stack = []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
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
			queue = temp
		}
	}
	return tree
}

func test01() {
	//          1
	//       /    \
	//   null      2
	//            /
	//           3
	fmt.Println("preorderTraversal:", preorderTraversal(buildBinaryTree([]int{1, -1, 2, 3}))) // -1代表null
	fmt.Println("preorderTraversal:", preorderTraversal(buildBinaryTree([]int{})))
	fmt.Println("preorderTraversal:", preorderTraversal(buildBinaryTree([]int{1})))
	fmt.Println("preorderTraversal:", preorderTraversal(buildBinaryTree([]int{1, 2})))
	fmt.Println("preorderTraversal:", preorderTraversal(buildBinaryTree([]int{1, -1, 2})))

	fmt.Println("preorderTraversal_1:", preorderTraversal_1(buildBinaryTree([]int{1, -1, 2, 3}))) // -1代表null
	fmt.Println("preorderTraversal_1:", preorderTraversal_1(buildBinaryTree([]int{})))
	fmt.Println("preorderTraversal_1:", preorderTraversal_1(buildBinaryTree([]int{1})))
	fmt.Println("preorderTraversal_1:", preorderTraversal_1(buildBinaryTree([]int{1, 2})))
	fmt.Println("preorderTraversal_1:", preorderTraversal_1(buildBinaryTree([]int{1, -1, 2})))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
