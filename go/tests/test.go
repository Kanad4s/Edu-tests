package main

import (
	"fmt"
)

func main() {
	a := []int{6, 6, 5, 5, 4, 1}
	// a := []int{6, 17, 5, 38, 19, 94, 44, 47, 10, 70, 41, 90, 65, 100, 87, 95, 32, 36, 76, 95, 61, 18, 60, 28, 22, 27, 52, 16, 19, 24, 47, 10, 28, 35, 56, 40, 81, 83, 88, 67, 66, 89, 8, 88, 16, 16, 44, 95}
	fmt.Println(majorityElement(a))
}
func majorityElement(nums []int) int {
	count := make(map[int]int)
	maxNum := 0
	maxCount := 0
	for _, num := range nums {
		count[num] += 1
		if count[num] > maxCount {
			maxCount = count[num]
			maxNum = num
		}
	}
	return maxNum
}
