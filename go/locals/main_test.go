package main

import (
	"testing"
)

func BenchmarkSlowSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		a := []int{12, 34, 4, 22, 45, 11, 32, 4, 99, 9, 2, 4, 1, 10, 111}
		SlowSort(a)
	}
}

func BenchmarkOptimizedSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{12, 34, 4, 22, 45, 11, 32, 4, 99, 9, 2, 4, 1, 10, 111, 100, 432, 5, 43}
		OptimizedSort(a)
	}
}

func BenchmarkInefficientStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []string{"12", "34", "4", "22", "45", "11", "32", "4", "99", "9", "2", "4", "", "10", "111", "100", "432"}
		InefficientStringBuilder(a, len(a))
	}
}

func BenchmarkOptimizedStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []string{"12", "34", "4", "22", "45", "11", "32", "4", "99", "9", "2", "4", "", "10", "111", "100", "432"}
		OptimizedStringBuilder(a, len(a))
	}
}

func BenchmarkExpensiveCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExpensiveCalculation(30)
	}
}

func BenchmarkOptimizedCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptimizedCalculation(30)
	}
}

func BenchmarkHighAllocationSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "AAaBBbbAababaabbbbbbfaa"
		subStr := "aa"
		HighAllocationSearch(str, subStr)
	}
}

func BenchmarkOptimizedSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "AAaBBbbAababaabbbbbbfaa"
		subStr := "aa"
		OptimizedSearch(str, subStr)
	}
}
