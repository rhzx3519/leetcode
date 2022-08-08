/**
@author ZhengHao Lou
Date    2022/07/11
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/implement-magic-dictionary/
*/

type MagicDictionary struct {
	dict map[string]bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{dict: map[string]bool{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		this.dict[s] = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for key := range this.dict {
		if len(key) == len(searchWord) {
			var c int
			for i := range key {
				if key[i] != searchWord[i] {
					c++
				}
				if c > 1 {
					break
				}
			}
			if c == 1 {
				return true
			}
		}
	}
	return false
}

func main() {
	d := Constructor()
	//["MagicDictionary", "buildDict", "search", "search", "search", "search"]
	//	[[], [["hello","hallo","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
	// [null,null,true,true,false,false]
	d.BuildDict([]string{"hello", "hallo", "leetcode"})
	fmt.Println(d.Search("hello"))
	fmt.Println(d.Search("hhllo"))
	fmt.Println(d.Search("hell"))
	fmt.Println(d.Search("leetcoded"))
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
