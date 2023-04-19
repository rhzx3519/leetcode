package main

/*
*
https://leetcode.cn/problems/maximum-score-words-formed-by-letters/
思路：状态压缩DP
*/
func maxScoreWords(words []string, letters []byte, score []int) (tot int) {
	n := len(words)
	cnt := make(map[int]int)
	for i := range letters {
		cnt[int(letters[i]-'a')]++
	}
	for mask := 1; mask < (1 << n); mask++ {
		wordCnt := make(map[int]int)
		for i := range words {
			if mask&(1<<i) == 0 {
				continue
			}
			for j := range words[i] {
				wordCnt[int(words[i][j]-'a')]++
			}
		}
		var s int
		var ok = true
		for j := 0; j < 26; j++ {
			s += score[j] * wordCnt[j]
			ok = ok && (wordCnt[j] <= cnt[j])
		}
		if ok && s > tot {
			tot = s
		}
	}
	return
}

func main() {
	maxScoreWords([]string{"dog", "cat", "dad", "good"},
		[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
		[]int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
