/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/rank-teams-by-votes/
 */
func rankTeams(votes []string) string {
	n := len(votes[0])
	mp := make(map[byte][]int)
	for _, c := range votes[0] {
		mp[byte(c)] = make([]int, n)
	}

	for _, vote := range votes {
		for i, c := range vote {
			mp[byte(c)][i]++
		}
	}

	keys := []byte(votes[0])
	sort.Slice(keys, func(i, j int) bool {
		a, b := keys[i], keys[j]
		for i := 0; i < n; i++ {
			if mp[a][i] != mp[b][i] {
				return mp[a][i] > mp[b][i]
			}
		}
		return  a < b
	})
	return string(keys)
}

func main() {
	fmt.Println(rankTeams([]string{"ABC","ACB","ABC","ACB","ACB"}))
	fmt.Println(rankTeams([]string{"WXYZ","XYZW"}))
	fmt.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}))
	fmt.Println(rankTeams([]string{"BCA","CAB","CBA","ABC","ACB","BAC"}))
	fmt.Println(rankTeams([]string{"M","M","M","M"}))
}
