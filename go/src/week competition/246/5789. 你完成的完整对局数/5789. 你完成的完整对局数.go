package main

import (
	"fmt"
	"strconv"
)

func numberOfRounds(startTime string, finishTime string) int {
	s1, _ := strconv.Atoi(startTime[:2])
	s2, _ := strconv.Atoi(startTime[3:])
	f1, _ := strconv.Atoi(finishTime[:2])
	f2, _ := strconv.Atoi(finishTime[3:])

	var hour int
	if f1 < s1 || (f1==s1 && s2 > f2) {
		hour = f1 + 24 - s1
	} else {
		hour = f1 - s1
	}
	var total = 4*hour
	total -= left(s2)
	total += right(f2)

	return total
}

func left(m int) int {
	if m == 0 {
		return 0
	} else if m <= 15 {
		return 1
	} else if m <= 30 {
		return 2
	} else if m <= 45 {
		return 3
	} else {
		return 4
	}
}
func right(m int) int {
	if m == 60 {
		return 4
	} else if m >= 45 {
		return 3
	} else if m >= 30 {
		return 2
	} else if m >= 15 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(numberOfRounds("12:01", "12:44"))
	fmt.Println(numberOfRounds("20:00", "06:00"))
	fmt.Println(numberOfRounds("00:00", "23:59"))
	fmt.Println(numberOfRounds("00:01", "00:00"))
}