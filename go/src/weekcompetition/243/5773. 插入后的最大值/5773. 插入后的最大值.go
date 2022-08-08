package main

import (
	"fmt"
	"strconv"
)

func maxValue(s string, x int) string {
	var d = true
	if s[0] == '-' {
		i := find2(s[1:], x) + 1
		return s[:i] + strconv.Itoa(x) + s[i:]
	}

	i := find(s, x, d)
	return s[:i] + strconv.Itoa(x) + s[i:]
}

func find(s string, x int, d bool) int {
	var i int
	for ; i < len(s) && int(s[i]-'0') >= x; i++ {

	}
	return i
}

func find2(s string, x int) int {
	var i int
	for ; i < len(s) && int(s[i]-'0') <= x; i++ {

	}
	return i
}

func main() {
	fmt.Println(maxValue("873", 6))
	fmt.Println(maxValue("-55", 2))
	fmt.Println(maxValue("99", 9))
	fmt.Println(maxValue("-13", 2))
	fmt.Println(maxValue("28824579515", 8))
	fmt.Println(maxValue("-132", 3))
	fmt.Println(maxValue("469975787943862651173569913153377", 3))
}
