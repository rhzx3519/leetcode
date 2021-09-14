package main

import "fmt"

func checkValidString(s string) bool {
	left := []int{}
	star := []int{}

	for i := range s {
		switch s[i] {
		case '(':
			left = append(left, i)
		case '*':
			star = append(star, i)
		case ')':
			if len(left) != 0 {
				left = left[:len(left)-1]
			} else if len(star) != 0 {
				star = star[:len(star)-1]
			} else {
				return false
			}
		}
	}

	for len(left) > 0 && len(star) > 0 {
		if left[len(left)-1] > star[len(star)-1] {
			return false
		}
		left = left[:len(left)-1]
		star = star[:len(star)-1]
	}

	return len(left) == 0
}

func main() {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(*)))"))
	fmt.Println(checkValidString("("))
	fmt.Println(checkValidString("(*"))
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
}
