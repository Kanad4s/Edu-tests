package main

import (
	"fmt"
)

func main() {
	a := make(map[[2]int]int)
	for i := 0; i < 10; i++ {
		a[[2]int{i, i * 10}] = i % 3
		// a[i] = []int{i, i * 10}
		// tmp := make([]int, 2)
		// tmp[0] = i
		// tmp[1] = i * 10
		// a[i] = append(a[i], tmp)
	}
	res := make(map[int]int)
	for k, v := range a {
		fmt.Printf("key: %d, value: %v\n", k, v)
		_, ok := res[v]
		if !ok {
			res[v] = 1
		} else {
			res[v]++
		}
	}
	fmt.Println(res)
}