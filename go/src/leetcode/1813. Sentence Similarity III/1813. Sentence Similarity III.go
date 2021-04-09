package main

import (
	"fmt"
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")
	if len(s1) > len(s2) {
		return areSentencesSimilar(sentence2, sentence1)
	}

	mapper1 := make(map[string]int)
	mapper2 := make(map[string]int)
	for _, w := range s1 {
		mapper1[w]++
	}
	for _, w := range s2 {
		mapper2[w]++
	}

	for w1, n1 := range mapper1 {
		n2, ok := mapper2[w1]
		if !ok || n1 > n2 {
			return false
		}
	}

	return traverse(mapper1, s2, true) || traverse(mapper1, s2, false)
}

func traverse(mapper map[string]int, s []string, reverse bool) bool {
	m := make(map[string]int)
	for k, v := range mapper {
		m[k] = v
	}

	n := len(s)
	if reverse {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			s[l], s[r] = s[r], s[l]
		}
	}
	last := -1
	count := 0
	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]--
			if m[s[i]] == 0 {
				delete(m, s[i])
			}
			last = i
		} else if last==i-1 {
			count++
		}

		if count >= 2 {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	//fmt.Println(areSentencesSimilar("of", "A lot of words"))
	//fmt.Println(areSentencesSimilar("Eating right now", "Eating"))
	//fmt.Println(areSentencesSimilar("Luky", "Lucccky"))
	//fmt.Println(areSentencesSimilar("My name is Haley", "name is Haley"))
	//fmt.Println(areSentencesSimilar("A", "a A b A b"))
	fmt.Println(areSentencesSimilar("jiRNY fW rN S bpL r T S nl vndZkF om oUm ilsf pvyNJObW F Uj QNJUek",
		"jiRNY fW rN S bpL r T Uj QNJUek"	))


}
