package main

/**
https://leetcode-cn.com/problems/delete-characters-to-make-fancy-string/
 */
func makeFancyString(s string) string {
	last := byte('0')
	var count int

	bytes := []byte{}
	for i := range s {
		if s[i] == last {
			count++
		} else {
			last = s[i]
			count = 1
		}
		if count < 3 {
			bytes = append(bytes, s[i])
		}
	}

	return string(bytes)
}
