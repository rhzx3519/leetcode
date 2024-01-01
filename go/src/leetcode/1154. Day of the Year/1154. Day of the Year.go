/*
*
@author ZhengHao Lou
Date    2021/12/21
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var days_of_month = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func dayOfYear(date string) int {
	var ans int
	dates := strings.Split(date, "-")
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])
	for i := 1; i < month; i++ {
		if i == 2 && isLeapYear(year) {
			ans += 29
			continue
		}
		ans += days_of_month[i]
	}

	ans += day
	fmt.Println(ans)
	return ans
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	dayOfYear("2019-01-09")
	dayOfYear("2019-02-10")
	dayOfYear("2003-03-01")
	dayOfYear("2004-03-01")
}
