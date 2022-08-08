/**
@author ZhengHao Lou
Date    2022/01/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/day-of-the-week/
 */
var (
	days_of_month = [13]int{0,31,28,31,30,31,30,31,31,30,31,30,31}
	week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
)

// 1970/1/1 Thur
func dayOfTheWeek(day int, month int, year int) string {
	d:= getDays(day, month, year)
	d = (d + 2) % 7

	fmt.Println(d)
	return week[d]
}

func getDays(day int, month int, year int) int {
	var c int
	for i := 1970; i < year; i++ {
		c += 365
		if isLeapYear(i) {
			c++
		}
	}
	for i := 1; i < month; i++ {
		c += days_of_month[i]
		if i== 2 && isLeapYear(year) {
			c += 1
		}
	}
	c += day
	return c
}


func isLeapYear(year int) bool {
	return year % 4 == 0 && !(year%100 == 0 && year%400 != 0)
}

func main() {
	dayOfTheWeek(31, 8, 2019)
	dayOfTheWeek(18, 7, 1999)
	dayOfTheWeek(15, 8, 1993)
	dayOfTheWeek(31, 12, 2021)	// 4
	dayOfTheWeek(6, 3, 1995)	// 0
	dayOfTheWeek(21, 12, 1980)	// 6
}