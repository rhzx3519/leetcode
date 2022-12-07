/**
@author ZhengHao Lou
Date    2022/11/17
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-matching-subsequences/
*/
const N = 26

type pair struct {
	i, j int
}

func numMatchingSubseq(s string, words []string) (tot int) {
	var ps = make([][]pair, N)
	for i, word := range words {
		j := word[0] - 'a'
		ps[j] = append(ps[j], pair{i, 0})
	}

	for i := range s {
		var tmp []pair
		for _, pr := range ps[s[i]-'a'] {
			if len(words[pr.i])-1 == pr.j {
				tot++
			} else {
				tmp = append(tmp, pair{pr.i, pr.j + 1})
			}
		}
		ps[s[i]-'a'] = nil

		for _, pr := range tmp {
			ps[words[pr.i][pr.j]-'a'] = append(ps[words[pr.i][pr.j]-'a'], pr)
		}
	}

	return
}

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Println(numMatchingSubseq("abcde", []string{"acd"}))
	fmt.Println(numMatchingSubseq("rwpddkvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvjubjgloeofnpjqlkdsqvruvabjrikfwronbrdyyjnakstqjac",
		[]string{"wpddkvbnn", "lnagtva", "kvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvju", "rwpddkvbnnugln", "gloeofnpjqlkdsqvruvabjrikfwronbrdyyj", "vbgeinupkvgmgxeaaiuiyojmoqkahwvbpwugdainxciedbdkos", "mspuhbykmmumtveoighlcgpcapzczomshiblnvhjzqjlfkpina", "rgmliajkiknongrofpugfgajedxicdhxinzjakwnifvxwlokip", "fhepktaipapyrbylskxddypwmuuxyoivcewzrdwwlrlhqwzikq", "qatithxifaaiwyszlkgoljzkkweqkjjzvymedvclfxwcezqebx"}))
}
