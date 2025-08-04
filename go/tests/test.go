package main

import (
	"fmt"
	"strings"
)

func main() {
	// a := []int{4, 2, 2, 4, 3, 3, 3, 3}
	// b := []int{5, 9, 9, 0, 9, 6, 9, 6, 9, 9, 9}
	fmt.Println(makeFancyString("aabbbbrrreweewwww"))
}

func makeFancyString(s string) string {
	res := strings.Builder{}
	var letter rune
	count := 0
	for _, cur := range s {
		if cur == letter {
			count++
		} else {
			letter = cur
			count = 0
		}
		if count < 2 {
			res.WriteRune(cur)
		}
	}
	return res.String()
}
