package main

import "strings"

func numOfStrings(patterns []string, word string) int {
	var count int
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			count++
		}
	}
	return count
}