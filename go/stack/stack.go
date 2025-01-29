package main

import (
	"fmt"
	"github.com/informitas/stack"
)

func main() {
	s := stack.NewStack[rune]()
	s.Push(')')
	fmt.Println(s.Top())
}