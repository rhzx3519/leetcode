/**
@author ZhengHao Lou
Date    2021/10/11
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/integer-to-english-words/
 */

var NUMBER_MAP = map[int]string{
	//0:	"Zero",
	1:	"One",
	2:	"Two",
	3:	"Three",
	4:	"Four",
	5:	"Five",
	6:	"Six",
	7:	"Seven",
	8:	"Eight",
	9:	"Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

var UNIT_MAP = map[int]string{
	0: "",
	1: "Thousand",
	2: "Million",
	3: "Billion",
}



func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	digits := []int{}
	for i := 0; num != 0; i++ {
		r := num % 10
		digits = append(digits, r)
		num /= 10
	}
	strs := []string{}
	for i := 0; i < len(digits); i += 3 {
		tmp := convert(digits[i:min(i+3, len(digits))])
		if len(tmp) == 0 {
			continue
		}
		if i/3 > 0 {
			tmp = append(tmp, UNIT_MAP[i/3])
		}
		strs = append(strs, tmp...)
		copy(strs[len(tmp):], strs[:])
		copy(strs[:len(tmp)], tmp)
	}
	fmt.Println(strings.Join(strs, " "))
	return strings.Join(strs, " ")
}

func convert(digits []int) []string {
	strs := []string{}
	if len(digits) == 0 {
		return strs
	}
	var  d0, d1, d2 int		// 个，十，百
	d0 = digits[0]
	if len(digits) >= 2 {
		d1 = digits[1]
	}
	if len(digits) >= 3 {
		d2 = digits[2]
	}
	if d2 != 0 {
		strs = append(strs, fmt.Sprintf("%v Hundred", NUMBER_MAP[d2]))
	}
	t := d1 * 10 + d0
	if t == 0 {
		return strs
	}

	if t < 20 {
		strs = append(strs, NUMBER_MAP[t])
	} else {
		strs = append(strs, NUMBER_MAP[d1*10])
		if d0 != 0 {
			strs = append(strs, NUMBER_MAP[d0])
		}
	}
	return strs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	numberToWords(1)
	//numberToWords(123)
	//numberToWords(12345)
	//numberToWords(1234567)
	//numberToWords(1234567891)
}