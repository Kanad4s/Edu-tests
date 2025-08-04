package main

import (
	"fmt"
)

func main() {
	// a := []int{4, 2, 2, 4, 3, 3, 3, 3}
	b := []int{5, 9, 9, 0, 9, 6, 9, 6, 9, 9, 9}
	fmt.Println(totalFruit(b))
}
func totalFruit(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}
	busketCountA, busketCountB := 1, 0
	busketTypeA, busketTypeB := fruits[0], 0
	maxCount := 0
	curCount := 1
	for i := 1; i < len(fruits); i++ {
		if fruits[i] != busketTypeA && fruits[i] != busketTypeB {
			if maxCount < curCount {
				maxCount = curCount
			}
			if fruits[i-1] == busketTypeA {
				curCount = busketCountA
				busketTypeB = fruits[i]
				busketCountB = 1
				busketCountA = 0
			} else if fruits[i-1] == busketTypeB {
				curCount = busketCountB
				busketTypeA = fruits[i]
				busketCountA = 1
				busketCountB = 0
			}
		} else if fruits[i] != busketTypeA || fruits[i] != busketTypeB {
			if fruits[i] == busketTypeB {
				busketCountA = 0
				busketCountB++
			} else {
				busketCountB = 0
				busketCountA++
			}
		} else {
			if fruits[i] == busketTypeA {
				busketCountA++
			} else if fruits[i] == busketTypeB {
				busketCountB++
			}
		}
		curCount++
		fmt.Printf("curCount: %d, typeA %d, CountA %d, typeB %d, countB: %d\n",
			curCount, busketTypeA, busketCountA, busketTypeB, busketCountB)
	}
	if curCount > maxCount {
		maxCount = curCount
	}
	return maxCount
}
