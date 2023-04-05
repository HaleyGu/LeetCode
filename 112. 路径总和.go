package main

import (
	"fmt"
)

// 112. 路径总和
// 简单

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。

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
// 时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
// 空间复杂度：O(H)，其中 H 是树的高度。空间复杂度主要取决于递归时栈空间的开销，最坏情况下，树呈现链状，空间复杂度为 O(N)。
// 平均情况下树的高度与节点数的对数正相关，空间复杂度为 O(logN)。
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var target = targetSum - root.Val
	if root.Left == nil && root.Right == nil && target == 0 {
		return true
	}
	return hasPathSum(root.Left, target) || hasPathSum(root.Right, target)
}

func buildBinaryTree(arr []int) *TreeNode {
	var arrLen = len(arr)
	if arrLen == 0 {
		return nil
	}

	var tree = &TreeNode{Val: arr[0]}
	if arrLen == 1 {
		return tree
	}

	var index = 1
	var queue = []*TreeNode{tree}
	for len(queue) > 0 {
		var temp []*TreeNode
		for i := range queue {
			if index < arrLen && arr[index] != -1 {
				queue[i].Left = &TreeNode{Val: arr[index]}
				temp = append(temp, queue[i].Left)
			}
			index++
			if index < arrLen && arr[index] != -1 {
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
	fmt.Println("hasPathSum:", hasPathSum(buildBinaryTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}), 22))
	fmt.Println("hasPathSum:", hasPathSum(buildBinaryTree([]int{1, 2, 3}), 5))
	fmt.Println("hasPathSum:", hasPathSum(buildBinaryTree([]int{2, 3, 1}), 5))
	fmt.Println("hasPathSum:", hasPathSum(buildBinaryTree([]int{}), 0))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
