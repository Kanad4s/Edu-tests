package main

import "fmt"

const (
	i         = 5
	f         = 1.5
	i64 int64 = 64
)

var vTest = 1

func div(a, b int) (int, error) {
	return a / b, nil
}

func main() {
	fmt.Println(f+float64(vTest), i64+int64(vTest))
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			fmt.Println("Hello, Goб чет")
		} else if i%3 == 0 && (i > 3 || i%2 == 1) {
			fmt.Println("Hello, Go, нечет, комплексное условие")
		}
	}
}
