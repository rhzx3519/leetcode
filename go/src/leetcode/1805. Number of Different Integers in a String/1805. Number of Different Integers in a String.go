package main

import (
	"fmt"
	"strings"
)

func numDifferentIntegers(word string) int {
	tmp := []rune{}
	for _, ch := range word {
		if ch < '0' || ch > '9' {
			tmp = append(tmp, ' ')
		} else {
			tmp = append(tmp, ch)
		}
	}
	strs := strings.Split(string(tmp), " ")
	count := make(map[string]int)
	for _, str := range strs {
		t := strings.TrimSpace(str)
		if len(t) == 0 {
			continue
		}
		i := 0
		for i < len(t)-1 && t[i] == '0' {
			i++
		}

		count[str[i:]]++
	}
	fmt.Println(len(count))
	return len(count)
}

func main() {
	numDifferentIntegers("uu717761265362523668772526716127260222200144937319826j717761265362523668772526716127260222200144937319826k717761265362523668772526716127260222200144937319826b7177612653625236687725267161272602222001449373198262hgb9utohfvfrxprqva3y5cdfdudf7zuh64mobfr6mhy17")
	//numDifferentIntegers("0a0")
}