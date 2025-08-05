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

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	count := 0
	for _, fruit := range fruits {
		for j, basket := range baskets {
			if fruit <= basket {
				baskets[j] = 0
				break
			}
			if j == len(baskets)-1 {
				count++
			}
		}
		// fmt.Println(baskets)
	}
	return count
}
