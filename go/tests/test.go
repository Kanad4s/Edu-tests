package main

import (
	"fmt"
)

func main() {
	// a := []int{4, 2, 2, 4, 3, 3, 3, 3}
	// b := []int{5, 9, 9, 0, 9, 6, 9, 6, 9, 9, 9}
	fmt.Println(makeFancyString("aabbbbrrreweewwww"))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return isPath(root, 0, targetSum)
}

func isPath(root *TreeNode, sum, target int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val+sum == target
	}
	return isPath(root.Left, sum+root.Val, target) || isPath(root.Right, sum+root.Val, target)
}
