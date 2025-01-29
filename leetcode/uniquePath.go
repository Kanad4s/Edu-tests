package main

import (
	"fmt"
)

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
    aboveRow := make([]int, n, n)
	for i, _ := range aboveRow {
		aboveRow[i] = 1
	}
    fmt.Println(aboveRow)
    for i := 1; i < m; i++ {
        currentRow := make([]int, n, n)
		for i, _ := range currentRow {
			currentRow[i] = 1
		}
        for j := 1; j < n; j++ {
            currentRow[j] = currentRow[j - 1] + aboveRow[j]
        }
        aboveRow = currentRow
    }

    return aboveRow[n - 1]
}
