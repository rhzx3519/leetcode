package main

import "fmt"

func minSwaps(s string) int {
	var count int
	bytes := []byte(s)
	n := len(s)
	var cl, cr int
	l, r := 0, n-1
	for l < r {

		for l < r && cl >= 0 {
			if bytes[l] == ']' {
				cl--
			} else {
				cl++
			}
			l++
		}
		if cl < 0 {
			l--
		}

		for l < r && cr >= 0 {
			if bytes[r] == '[' {
				cr--
			} else {
				cr++
			}
			r--
		}
		if cr < 0 {
			r++
		}

		if l < r && cl < 0 && cr < 0 {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			count++
			cl++
			cr++
		}
	}
	fmt.Println(string(bytes))
	fmt.Println(cl, cr)
	fmt.Println(count)
	return count
}

func main() {
	//minSwaps("[[[]]]][][]][[]]][[[") // [[][]][[]]][[]	// [[]]][
	//minSwaps("][][")
	//minSwaps("]]][[[")
	minSwaps("[]")
}
