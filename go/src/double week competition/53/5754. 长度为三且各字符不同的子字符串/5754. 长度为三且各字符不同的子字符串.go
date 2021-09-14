package main

import "fmt"

func countGoodSubstrings(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}
	var count int
	mapper := make(map[byte]int)
	for i := 0; i < n; i++ {
		mapper[s[i]]++
		if i > 2 {
			mapper[s[i-3]]--
		}

		if i >= 2 {
			var f bool
			for _, cnt := range mapper {
				if cnt > 1 {
					f = true
					break
				}
			}

			if !f {
				count++
			}
		}


	}

	return count
}

func main() {
	fmt.Println(countGoodSubstrings("xyzzaz"))
	fmt.Println(countGoodSubstrings("aababcabc"))
}
