/**
@author ZhengHao Lou
Date    2022/01/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-all-possible-recipes-from-given-supplies/
思路：拓扑排序
 */
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	var ans = []string{}
	sup := make(map[string]bool)
	for _, s := range supplies {
		sup[s] = true
	}

	graph := make(map[string][]string)
	for _, recipe := range recipes {
		graph[recipe] = []string{}
	}

	que := []string{}
	ind := make(map[string]int)
	for i, recipe := range recipes {
		for _, ig := range ingredients[i] {
			if sup[ig] == false {
				graph[ig] = append(graph[ig], recipe)
				ind[recipe]++
			}
		}
		if ind[recipe] == 0 {
			que = append(que, recipe)
		}
	}

	for len(que) > 0 {
		r := que[0]
		que = que[1:]
		ans = append(ans, r)
		for _, nr := range graph[r] {
			ind[nr]--
			if ind[nr] == 0 {
				que = append(que, nr)
			}
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	findAllRecipes([]string{"bread"}, [][]string{{"yeast","flour"}}, []string{"yeast","flour","corn"})
}