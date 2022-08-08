/**
@author ZhengHao Lou
Date    2021/10/19
*/
package main

/**
https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/
 */
type WordDictionary struct {
	dict map[int]map[string]int
}


func Constructor() WordDictionary {
	return WordDictionary{
		dict: map[int]map[string]int{},
	}
}


func (this *WordDictionary) AddWord(word string)  {
	n := len(word)
	if _, ok := this.dict[n]; !ok {
		this.dict[n] = map[string]int{}
	}
	this.dict[n][word]++
}


func (this *WordDictionary) Search(word string) bool {
	n := len(word)
	if _, ok := this.dict[n]; !ok {
		return false
	}
	var mapper map[string]int = this.dict[n]
	for w, _ := range mapper {
		if match(word, w) {
			return true
		}
	}
	return false
}

func match(a, b string) bool {
	for i := range a {
		if a[i] != '.' && a[i] != b[i] {
			return false
		}
	}
	return true
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
