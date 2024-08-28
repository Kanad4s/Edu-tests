package main

import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			fmt.Println("Hello, Goб чет")
		} else if i%3 == 0 && (i > 3 || i%2 == 1) {
			fmt.Println("Hello, Go, нечет, комплексное условие")
		}
	}
}
