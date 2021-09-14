package main

import "fmt"

func maxLength(arr []string) int {
	n := len(arr)

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i == n {
			return 0
		}
		var t = m
		for _, c := range arr[i] {
			if m>>(c-'a') & 1 == 1 {
				return dfs(i+1, t)
			}
			m |= 1<<(c-'a')
		}
		var l = len(arr[i])
		return max(l + dfs(i+1, m), dfs(i+1, t))
	}

	return dfs(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxLength([]string{"un","iq","ue"}))
	fmt.Println(maxLength([]string{"cha","r","act","ers"}))
	fmt.Println(maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
	fmt.Println(maxLength([]string{"jnfbyktlrqumowxd","mvhgcpxnjzrdei"}))
}
