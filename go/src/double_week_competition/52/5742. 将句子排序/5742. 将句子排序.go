package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	sort.Slice(words, func(i, j int) bool {
		return words[i][len(words[i])-1] < words[j][len(words[j])-1]
	})

	newS := []string{}
	for _, w := range words {
		newS = append(newS, w[:len(w)-1])
	}
	return strings.Join(newS, " ")
}

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}