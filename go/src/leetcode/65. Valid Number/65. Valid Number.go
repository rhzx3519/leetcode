package main

import "fmt"

func isNumber(s string) bool {
	n := len(s)
	var i int
	for i < n && (s[i] != 'e' && s[i] != 'E') {
		i++
	}
	if i < n {
		return isDecimalOrInteger(s[:i]) && isInteger(s[i+1:])
	}
	return isDecimalOrInteger(s)
}

func isDecimalOrInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	n := len(s)
	var i int
	for i < n && isBlank(s[i]) {
		i++
	}
	if i == n {
		return false
	}

	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		i++
	}
	if i == n {
		return false
	}
	if s[i] == '.' {
		i++
		var j = i
		for j < n && isDigit(s[j]) {
			j++
		}
		if i == j {
			return false
		}
		return j == n
	} else if isDigit(s[i]) {
		for i < n && isDigit(s[i]) {
			i++
		}
		if i == n {
			return true
		}
		if s[i] == '.' {
			i++
		}
		for i < n && isDigit(s[i]) {
			i++
		}
		return i == n
	}
	return false
}

func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	n := len(s)
	var i int
	if isSign(s[i]) {
		i++
	}
	if i == n {
		return false
	}
	for i < n && isDigit(s[i]) {
		i++
	}
	return i == n
}

func isBlank(c byte) bool {
	return c == ' '
}

func isDigit(c byte) bool {
	return c >= '0' && c <='9'
}

func isSign(c byte) bool {
	return c == '+' || c == '-'
}

func main() {
	for i, s := range []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10",
		"-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"} {
		fmt.Println(i, isNumber(s))
	}

	fmt.Println("---------------------")

	for i, s := range []string{".", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"} {
		fmt.Println(i, isNumber(s))
	}
}
