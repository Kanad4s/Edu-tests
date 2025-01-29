package main

import (
	"fmt"
)

// func main() {
// 	a := 7
// 	fmt.Println(climbStairs(a))
// 	fmt.Println(climbStairs1(a))
// 	fmt.Println(intToRoman(2813))
// 	fmt.Println(merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3))
// 	// fmt.Println(merge([]int{0}, 0, []int{1}, 1))
// }

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i := 0
	j := smallest(nums2, n)
	for n > 0 && nums2[n - 1] != 0 {
		if i < m && j < n && nums1[i] < nums2[j] {
			fmt.Printf("nums1[i]: %d, nums2[j]: %d\n", nums1[i], nums2[j])
			i++
		} else {
			swap(&nums1[i], &nums2[j])
			i++
			j = smallest(nums2, n)
			fmt.Printf("new j: %d\n", j)                                 
			fmt.Println(nums2)
		}
	}
	return nums1
}

func smallest(nums []int, n int) int {
    index := n - 1
    for i, _ := range nums {
        if nums[i] < nums[index] && nums[i] != 0 {
            index = i
        }
    }
    return index
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func intToRoman(num int) string {
	/*M-1000, D-500, C-100, L-50, X-10, V-5, I-1*/
	digit := num / 1000
	roman := make([]byte, 0, 1)
	for i := 0; i < digit; i++ {
		roman = append(roman, 'M')
	}
	num %= 1000
	digit = num / 100
	insertRomanDigit(&roman, digit, 'M', 'D', 'C')
	num %= 100
	digit = num / 10
	insertRomanDigit(&roman, digit, 'C', 'L', 'X')
	num %= 10
	digit = num
	insertRomanDigit(&roman, digit, 'X', 'V', 'I')
	return string(roman)
}

func insertRomanDigit(roman *[]byte, digit int, oldSymbol byte, mediumSymbol byte, smallSymbol byte) {
	if digit == 9 {
		*roman = append(*roman, smallSymbol)
		*roman = append(*roman, oldSymbol)
	} else if digit >= 5 {
		*roman = append(*roman, mediumSymbol)
		left := digit - 5
		for i := 0; i < left; i++ {
			*roman = append(*roman, smallSymbol)
		}
	} else if digit == 4 {
		*roman = append(*roman, smallSymbol)
		*roman = append(*roman, mediumSymbol)
	} else {
		for i := 0; i < digit; i++ {
			*roman = append(*roman, smallSymbol)
		}
	}
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}

func climbStairs1(n int) int {
	count := n / 2
	sum := 0
	for i := 0; i <= count; i++ {
		sum += combination(n-i, i)
	}
	return sum
}

func combination(n int, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

/*
2 - 2
3 - 3
4 - 5
5 - 8
6 - 13
7 - 21


1,1,1,1...
2,2,2,1
2,2,1,2
2,1,2,2
1,2,2,2

1,2,1
2,1,1
1,1,2



n
2,2,1,1,1
2,1,2,1,1
2,1,1,2,1
2,1,1,1,2

6
*/
