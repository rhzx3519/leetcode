/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

/**
https://leetcode-cn.com/problems/split-two-strings-to-make-palindrome/
 */
func checkPalindromeFormation(a string, b string) bool {
	var la, ra = center(a)
	var lb, rb = center(b)

	return check(a, b, la, rb) || check(b, a, lb, ra)
}

func center(a string) (int, int) {
	n := len(a)
	var la, ra int
	for la, ra = (n-1)>>1, n>>1; la >= 0 && ra < n && a[la] == a[ra]; la, ra = la-1, ra+1 {}
	return la, ra
}

func check(a string, b string, la int, rb int) bool {
	n := len(a)
	var i int
	for i < n && a[i] == b[n-1-i] {
		i++
	}
	return i > la || n-1-i < rb
}

func main() {
	checkPalindromeFormation("a", "a")
	checkPalindromeFormation("x", "y")
	checkPalindromeFormation("abdef", "fecab")
	checkPalindromeFormation("ulacfd", "jizalu")
	checkPalindromeFormation("xbdef", "xecab")	// false
	checkPalindromeFormation("abda", "acmc")		// false
}

//"pvhmupgqeltozftlmfjjde"
//"yjgpzbezspnnpszebzmhvp"
