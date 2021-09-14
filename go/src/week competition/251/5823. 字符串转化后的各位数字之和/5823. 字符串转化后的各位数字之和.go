package main

import "fmt"

func getLucky(s string, k int) int {
	var sum int
	for _, c := range s {
		t := int(c - 'a') + 1
		if t < 10 {
			sum += t
		} else {
			sum += t/10
			sum += t%10
		}
	}
	k--
	for k > 0 {
		k--
		tmp := 0
		for sum != 0 {
			tmp += sum%10
			sum /= 10
		}
		sum = tmp
	}
	return sum
}

func main() {
	fmt.Println(getLucky("iiii", 1))
	fmt.Println(getLucky("leetcode", 2))
}
