/**
@author ZhengHao Lou
Date    2022/05/30
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
https://leetcode.cn/problems/apply-discount-to-prices/
*/
func discountPrices(sentence string, discount int) string {
	var words []string
	for _, w := range strings.Split(sentence, " ") {
		if is, p := isPrice(w); is {
			words = append(words, fmt.Sprintf("$%.2f", p*float64(100-discount)/float64(100)))
		} else {
			words = append(words, w)
		}
	}
	return strings.Join(words, " ")
}

func isPrice(w string) (bool, float64) {
	n := len(w)
	if n <= 1 {
		return false, 0
	}
	var i int
	if w[i] != '$' {
		return false, 0
	}
	i++
	var p float64
	for ; i < n && unicode.IsDigit(rune(w[i])); i++ {
		p = p*10 + float64(w[i]-'0')
	}
	if i == n {
		return true, p
	}
	if w[i] != '.' {
		return false, 0
	}
	i++
	if i == n {
		return true, p
	}

	var l = 1
	var decimals float64
	for ; i < n && unicode.IsDigit(rune(w[i])); i++ {
		decimals = decimals*10 + float64(w[i]-'0')
		l *= 10
	}
	if i != n || l == 1 {
		return false, 0
	}
	decimals /= float64(l)
	return true, p + decimals
}

func main() {
	//fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 50))
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))
}
