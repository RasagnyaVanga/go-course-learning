package main

import (
	"fmt"
	"regexp"
	"strings"
)

func WordCount(input string) map[string]int {
	freq := make(map[string]int)
	wordPattern := regexp.MustCompile(`\b[\w']+\b`)
	words := wordPattern.FindAllString(input, -1)

	for _, word := range words {
		lowercase := strings.ToLower(word)
		freq[lowercase]++
	}

	return freq
}
func main() {
	text := "You come back, you hear me? DO YOU HEAR ME?"
	result := WordCount(text)

	for word, count := range result {
		fmt.Printf("%s: %d\n", word, count)
	}
}
