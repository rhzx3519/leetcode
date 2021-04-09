package main

import "fmt"

var mp = map[uint8]string {
	'2' : "abc",
	'3' : "def",
	'4' : "ghi",
	'5' : "jkl",
	'6' : "mno",
	'7' : "pqrs",
	'8' : "tuv",
	'9' : "wxyz",
}

func letterCombinations(digits string) []string {
	ans := []string{}
	dfs(digits, 0, &[]uint8{}, &ans)
	return ans
}

func dfs(digits string, idx int, arr *[]uint8, ans *[]string)  {
	if len(*arr) > len(digits) {
		return
	}

	if idx == len(digits) {
		if len(*arr) != 0 && len(*arr) == len(digits) {
			*ans = append(*ans, string(*arr))
		}
		return
	}

	for i := idx; i < len(digits); i++ {
		for _, c := range mp[digits[i]] {
			*arr = append(*arr, uint8(c))
			dfs(digits, i+1, arr, ans)
			*arr = (*arr)[:len(*arr)-1]
		}
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
