/**
@author ZhengHao Lou
Date    2022/01/09
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-words-obtained-after-adding-a-letter/
思路：位运算
具体的，将字符串转换为二进制存储, 遍历去掉target中的一位，查看是否存在为startWords
time: O(n)
space: O(n)
n = max(len(startWords), len(targetWords))
 */
func wordCount(startWords []string, targetWords []string) int {
	var ans int
	exists := make(map[int]bool)
	for _, w := range startWords {
		exists[compress(w)] = true
	}
	for _, w := range targetWords {
		mask := compress(w)
		for i := range w {
			r := int(w[i] - 'a')
			if exists[mask - 1<<r] {
				ans++
				break
			}
		}
	}

	fmt.Println(ans)
	return ans
}

func compress(w string) int {
	var r int
	for i := range w {
		r |= 1 << int(w[i] - 'a')
	}
	return r
}

func main() {
	wordCount([]string{"ant","act","tack"}, []string{"tack","act","acti"})
	wordCount([]string{"ab","a"}, []string{"abc","abcd"})
}