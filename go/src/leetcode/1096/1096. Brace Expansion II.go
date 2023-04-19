package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
*
https://leetcode.cn/problems/brace-expansion-ii/
思路：将表达式拆分为3部分
1. 从左向右寻找第一个出现的'}'出现的下标j，如果找不到，说明是字母直接返回;
2. 从右向左寻找expr[:j]出现的第一个'{'的下标i, 此时可以认为expr[i:j+1] {}表示一个完整的集合b
3. expr[:i], expr[j+1:]表示和b相邻的另外两个集合a, c
4. b根据','拆分，分别与另外两个集合a, c做一个拼接，之后通过递归求解a+b+c这个表达式的值
递归过程中使用map记录每一个求解结果
*/
func braceExpansionII(expression string) (ans []string) {
	var st = make(map[string]int)
	var dfs func(string)
	dfs = func(expr string) {
		j := strings.IndexByte(expr, '}')
		if j == -1 {
			st[expr]++
			return
		}
		i := strings.LastIndexByte(expr[:j], '{')
		a, c := expr[:i], expr[j+1:]
		for _, b := range strings.Split(expr[i+1:j], ",") {
			dfs(a + b + c)
		}
	}

	dfs(expression)
	for k := range st {
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return
}

func main() {
	fmt.Println(braceExpansionII("{a,b}{c,{d,e}}"))
	fmt.Println(braceExpansionII("{{a,z},a{b,c},{ab,z}}"))
}
