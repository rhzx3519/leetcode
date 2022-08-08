/**
@author ZhengHao Lou
Date    2021/11/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/
 */
const N = 26
var (
	digits = []string{
			"zero",	//
			"one",	//
			"two",	//
			"three", //
			"four", //
			"five",	//
			"six",	//
			"seven",
			"eight", //
			"nine",
		}
)

func originalDigits(s string) string {
	alphabet := make([]int, N)
	for i := range s {
		alphabet[index(s[i])]++
	}
	nums := make([]int, 10)
	if alphabet[index('z')] !=  0 {
		c := alphabet[index('z')]
		nums[0] += c
		for _, ch := range digits[0] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('x')] !=  0 {
		c := alphabet[index('x')]
		nums[6] += c
		for _, ch := range digits[6] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('w')] !=  0 {
		c := alphabet[index('w')]
		nums[2] += c
		for _, ch := range digits[2] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('g')] !=  0 {
		c := alphabet[index('g')]
		nums[8] += c
		for _, ch := range digits[8] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('u')] !=  0 {
		c := alphabet[index('u')]
		nums[4] += c
		for _, ch := range digits[4] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('o')] !=  0 {
		c := alphabet[index('o')]
		nums[1] += c
		for _, ch := range digits[1] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('r')] !=  0 {
		c := alphabet[index('r')]
		nums[3] += c
		for _, ch := range digits[3] {
			alphabet[index(byte(ch))] -= c
		}
	}

	if alphabet[index('f')] !=  0 {
		c := alphabet[index('f')]
		nums[5] += c
		for _, ch := range digits[5] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('s')] !=  0 {
		c := alphabet[index('s')]
		nums[7] += c
		for _, ch := range digits[7] {
			alphabet[index(byte(ch))] -= c
		}
	}
	if alphabet[index('i')] !=  0 {
		c := alphabet[index('i')]
		nums[9] += c
		for _, ch := range digits[9] {
			alphabet[index(byte(ch))] -= c
		}
	}

	bytes := []byte{}
	for i, c := range nums {
		for j := 0; j < c; j++ {
			bytes = append(bytes, byte('0' + i))
		}
	}
	return string(bytes)
}

func index(b byte) int {
	return int(b - 'a')
}

func main() {
	//fmt.Println(originalDigits("owoztneoer"))
	//fmt.Println(originalDigits("fviefuro"))
	fmt.Println(originalDigits("nnei"))
}