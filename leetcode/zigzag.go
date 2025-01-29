package main

import (
	"fmt"
)

// func main() {
// 	//    PINALSIGYAHRPI
// 	s := "PAYPALISHIRING"
// 	//    PINAPSIGYYHHPI
// 	rows := 4
// 	fmt.Println(convert(s, rows))
// }

func convert(s string, numRows int) string {
    res := ""
    length := len(s)
    if numRows == 1 || length <= 2 {
        return s
    }
    numColumns := len(s) / (numRows * 2 - 2)
    if len(s) % (numRows * 2 - 2) != 0 {
        numColumns++
    }
	letsInColumn := numRows * 2 - 2
	fmt.Println(numColumns)
	fmt.Println(letsInColumn)
    for i := 0; i < numRows; i++ {
        if i == 0 || i == numRows - 1 {
            // boundary
            for j := 0; j < numColumns; j++ {
                index := j * letsInColumn + i
                if index / length >= 1 {
                    break
                }
                res += string(s[index])
            }
            continue
        }
		fmt.Printf("i: %d inner circle, s: %s\n", i, res)
        for j := 0; j < numColumns; j++ {
            index := j * letsInColumn + i
			// fmt.Printf("index: %d\n", index)
            if index / length >= 1 {
                break
            }
            res += string(s[index])
            index = j * letsInColumn + i + (numRows - i - 1 + numRows - i - 1)
			// fmt.Printf("index: %d\n", index)
            if index / length >= 1 {
                break
            }
            res += string(s[index])
        }
		// fmt.Println(res)
    }
    return res
}