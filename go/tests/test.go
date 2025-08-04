package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{4, 4}
	b := []int{3, 3}
	fmt.Println(minCost1(a, b))
}

func minCost1(basket1 []int, basket2 []int) int64 {
	b1 := make(map[int]int)
	b2 := make(map[int]int)
	var cost int
	for _, fruit := range basket1 {
		b1[fruit]++
	}
	for _, fruit := range basket2 {
		b2[fruit]++
	}
	for k, v := range b1 {
		if (v+b2[k])%2 == 1 {
			return -1
		}
		cost += abs(v-b2[k]) / 2
	}
	for k, v := range b2 {
		_, ok := b1[k]
		if ok {
			continue
		}
		if v%2 == 1 {
			return -1
		}
		cost += v / 2
	}
	return int64(cost)
}

func minCost2(A, B []int) int64 {
	freq := map[int]int{}
	minVal := math.MaxInt
	for _, x := range A {
		freq[x]++
		if x < minVal {
			minVal = x
		}
	}
	for _, x := range B {
		freq[x]--
		if x < minVal {
			minVal = x
		}
	}

	extra := []int{}
	for k, v := range freq {
		if v%2 != 0 {
			return -1
		}
		for i := 0; i < abs(v)/2; i++ {
			extra = append(extra, k)
		}
	}

	sort.Ints(extra)
	var cost int64
	for i := 0; i < len(extra)/2; i++ {
		cost += int64(min(extra[i], 2*minVal))
	}
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
