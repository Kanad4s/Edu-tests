package main

import (
	"fmt"
)

func main() {
	a := []int{6, 6, 5, 5, 4, 1}
	// a := []int{6, 17, 5, 38, 19, 94, 44, 47, 10, 70, 41, 90, 65, 100, 87, 95, 32, 36, 76, 95, 61, 18, 60, 28, 22, 27, 52, 16, 19, 24, 47, 10, 28, 35, 56, 40, 81, 83, 88, 67, 66, 89, 8, 88, 16, 16, 44, 95}
	fmt.Println(countHillValley(a))
}

func countHillValley(nums []int) int {
	count := 0
	left := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			if (nums[i] > nums[left] && nums[i] > nums[i+1]) || (nums[i] < nums[left] && nums[i] < nums[i+1]) {
				count++
			}
			left = i
		}
	}
	return count
}
