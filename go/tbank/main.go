package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// var text string
	// text := "The quick brown fox jumps over the lazy dog."
	var text2 string
	text2 = "Go, go, go! Let's learn Go programming."
	mapa := CountWordFrequency(text2)
	fmt.Println(mapa)
}

func CountWordFrequency(text string) map[string]int {
	wordFreq := map[string]int{}
	text = strings.ToLower(text)
	// startIdx := 0
	isWord := false
	var word strings.Builder
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == '\'' {
			if char == '\'' {
				continue
			}
			if !isWord {
				isWord = true
			}
			word.WriteRune(char)
		} else {
			if isWord {
				wordFreq[word.String()]++
			}
			isWord = false
			word = strings.Builder{}
		}
		// fmt.Printf("%#U, %d\n", char, idx)
	}
	if isWord {
		wordFreq[word.String()]++
	}
	return wordFreq
}
