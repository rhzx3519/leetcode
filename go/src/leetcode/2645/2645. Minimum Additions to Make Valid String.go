package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-additions-to-make-valid-string/
*/
const abc = "abc"

func addMinimum(word string) (tot int) {
	n := len(word)
	for i := 0; i < n; {
		var k int
		for k < 3 && i < n {
			if word[i] == abc[k] {
				i++
				k++
			} else {
				tot++
				k++
			}
		}
		if k < 3 {
			tot += 3 - k
		}
	}

	return
}

func main() {
	fmt.Println(addMinimum("ab"))
}
