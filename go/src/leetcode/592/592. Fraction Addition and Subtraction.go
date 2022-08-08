/**
@author ZhengHao Lou
Date    2022/07/27
*/
package main

import (
	"fmt"
	"unicode"
)

/**
https://leetcode.cn/problems/fraction-addition-and-subtraction/
*/
type Fraction struct {
	Sign int
	A, B int
}

func (f *Fraction) Plus(other *Fraction) {
	a := f.A*other.B*f.Sign + other.A*f.B*other.Sign
	b := f.B * other.B
	if a < 0 {
		f.Sign = -1
		a = -a
	} else {
		f.Sign = 1
	}
	g := gcd(a, b)
	f.A = a / g
	f.B = b / g
}

func fractionAddition(expression string) string {
	var st []*Fraction
	n := len(expression)
	var j int
	for i := 0; i < len(expression); i = j {
		var frac = &Fraction{
			Sign: 1,
		}
		j = i
		if expression[j] == '+' {
			frac.Sign = 1
			j++
		} else if expression[j] == '-' {
			frac.Sign = -1
			j++
		}
		for j < n && unicode.IsDigit(rune(expression[j])) {
			frac.A = frac.A*10 + int(expression[j]-'0')
			j++
		}
		j++ // skip '/'
		for j < n && unicode.IsDigit(rune(expression[j])) {
			frac.B = frac.B*10 + int(expression[j]-'0')
			j++
		}
		
		if len(st) > 0 {
			other := st[len(st)-1]
			st = st[:len(st)-1]
			frac.Plus(other)
		}
		st = append(st, frac)
	}
	
	frac := st[len(st)-1]
	s := fmt.Sprintf("%v/%v", frac.A, frac.B)
	if frac.Sign == 1 {
		return s
	}
	return fmt.Sprintf("-%v", s)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	//fmt.Println(fractionAddition("-1/2+1/2"))
	//fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	//fmt.Println(fractionAddition("1/3-1/2"))
	fmt.Println(fractionAddition("7/2+2/3-3/4"))
}
