package main

import (
	"fmt"
	"slices"
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	makePath(root, 0, targetSum, &res, cur)
	return res
}

func makePath(root *TreeNode, sum, target int, res *[][]int, cur []int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val+sum == target {
			*res = append(*res, cur)
		}
		return
	}
	makePath(root.Left, sum+root.Val, target, res, cur)
	makePath(root.Right, sum+root.Val, target, res, slices.Clone(cur))
}
