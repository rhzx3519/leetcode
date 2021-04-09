package main

import "fmt"

func removeDuplicates(S string) string {
	st := []rune{}
	for _, c := range S {
		st = append(st, c)
		for len(st) >= 2 {
			if st[len(st)-1] != st[len(st)-2] {
				break
			}
			st = st[:len(st)-2]
		}
	}
	return string(st)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
