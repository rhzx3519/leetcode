package main

import "fmt"

func evaluate(s string, knowledge [][]string) string {
	kv := make(map[string]string)
	for _, t := range knowledge {
		kv[t[0]] = t[1]
	}

	tmp := []rune{}
	last := -1
	for i, ch := range s {
		if ch == '(' {
			last = i
		} else if ch == ')' {
			key := s[last+1:i]
			val, ok := kv[key]
			if ok {
				for _, t := range val {
					tmp = append(tmp, t)
				}
			} else {
				tmp = append(tmp, '?')
			}
			last = -1
		} else if last == -1 {
			tmp = append(tmp, ch)
		}
	}
	fmt.Println(string(tmp))
	return string(tmp)
}

func main() {
	evaluate("(name)is(age)yearsold", [][]string{{"name","bob"}, {"age","two"}})
	evaluate("hi(name)", [][]string{{"a","b"}})
	evaluate("(a)(a)(a)aaa", [][]string{{"a","yes"}})
}
