package main

import "fmt"

var N = 26

func minTimeToType(word string) int {
	var timecost int
	var point int
	for i := range word {
		c := word[i]
		target := int(c - 'a')
		d := (target + N - point) % N
		if d <= 13 {
			timecost += d + 1
		} else {
			timecost += N - d + 1
		}
		point = target
	}
	return timecost
}

func main() {
	fmt.Println(minTimeToType("abc"))
	fmt.Println(minTimeToType("bza"))
	fmt.Println(minTimeToType("zjpc"))
}
