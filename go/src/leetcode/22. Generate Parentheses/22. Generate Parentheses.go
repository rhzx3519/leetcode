package main

func generateParenthesis(n int) []string {
	ans := []string{}
	dfs(n, n, []uint8{}, &ans)
	return ans
}

func dfs(l, r int, arr []uint8, ans *[]string) {
	if l > r {
		return
	}
	if l==r && l==0 {
		*ans = append(*ans, string(arr))
		return
	}

	if l > 0 {
		dfs(l-1, r, append(arr, '('), ans)
	}
	if r > 0 {
		dfs(l, r-1, append(arr, ')'), ans)
	}
}
