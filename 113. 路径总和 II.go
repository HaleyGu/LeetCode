package main

import (
	"fmt"
)

// 113. 路径总和 II
// 中等

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

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

// 深度优先搜索
// 时间复杂度：O(N2)，其中 N 是树的节点数。
// 在最坏情况下，树的上半部分为链状，下半部分为完全二叉树，并且从根节点到每一个叶子节点的路径都符合题目要求。
// 此时，路径的数目为 O(N)，并且每一条路径的节点个数也为 O(N)，因此要将这些路径全部添加进答案中，时间复杂度为 O(N2)。
// 空间复杂度：O(N)，其中 N 是树的节点数。空间复杂度主要取决于栈空间的开销，栈中的元素个数不会超过树的节点数。
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ret [][]int
	var path []int
	var dfs func(root *TreeNode, targetSum int)

	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()
		var target = targetSum - root.Val
		if target == 0 && root.Left == nil && root.Right == nil {
			ret = append(ret, append([]int{}, path...))
		}

		dfs(root.Left, target)
		dfs(root.Right, target)
	}

	dfs(root, targetSum)
	return ret
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
	fmt.Println("pathSum:", pathSum(buildBinaryTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}), 22))
	fmt.Println("pathSum:", pathSum(buildBinaryTree([]int{2, 3, 1}), 5))
	fmt.Println("pathSum:", pathSum(buildBinaryTree([]int{1, 2, 3}), 5))
	fmt.Println("pathSum:", pathSum(buildBinaryTree([]int{1, 2}), 0))
	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
