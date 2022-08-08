package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/sentence-similarity-iii/
思路：
分成4种情况讨论:
1. sentence1 == sentence2
2. sentence1前缀一部分 + 后缀一部分 == sentence2
3. sentence1前缀一部分 == sentence2
4. sentence1后缀一部分 == sentence2

作者：answerer
链接：https://leetcode-cn.com/problems/sentence-similarity-iii/solution/shuang-duan-dui-lie-shi-xian-by-answerer-32as/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1, s2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return areSentencesSimilar(sentence2, sentence1)
	}
	// n1 <= n2
	for len(s1) != 0 && len(s2) != 0 {
		if s1[0] == s2[0] {
			copy(s1, s1[1:])
			s1 = s1[:len(s1) - 1]
			copy(s2, s2[1:])
			s2 = s2[:len(s2) - 1]
		} else if s1[len(s1) - 1] == s2[len(s2) - 1] {
			s1 = s1[:len(s1) - 1]
			s2 = s2[:len(s2) - 1]
		} else {
			break
		}
	}
	return len(s1) == 0
}

func main() {
	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	fmt.Println(areSentencesSimilar("of", "A lot of words"))
	fmt.Println(areSentencesSimilar("Eating right now", "Eating"))
	fmt.Println(areSentencesSimilar("Luky", "Lucccky"))
	fmt.Println(areSentencesSimilar("My name is Haley", "name is Haley"))
	fmt.Println(areSentencesSimilar("A", "a A b A b"))
	fmt.Println(areSentencesSimilar("jiRNY fW rN S bpL r T S nl vndZkF om oUm ilsf pvyNJObW F Uj QNJUek",
		"jiRNY fW rN S bpL r T Uj QNJUek"	))

}
