/**
@author ZhengHao Lou
Date    2022/05/31
*/
package main

/**
https://leetcode.cn/problems/Jf1JuT/
思路：拓扑排序
*/
func alienOrder(words []string) string {
	n := len(words)
	adj := make(map[byte][]byte)
	ind := make(map[byte]int)
	for i := range words[0] {
		ind[words[0][i]] = 0
	}

NEXT:
	for i := 1; i < n; i++ {
		for j := range words[i] {
			ind[words[i][j]] += 0
		}
		for j := 0; j < min(len(words[i-1]), len(words[i])); j++ {
			if words[i-1][j] != words[i][j] {
				adj[words[i-1][j]] = append(adj[words[i-1][j]], words[i][j])
				ind[words[i][j]]++
				continue NEXT
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return ""
		}
	}

	var que []byte
	for i := range ind {
		if ind[i] == 0 {
			que = append(que, i)
		}
	}

	var bytes []byte
	for len(que) != 0 {
		i := que[0]
		que = que[1:]
		bytes = append(bytes, i)
		for _, ni := range adj[i] {
			ind[ni]--
			if ind[ni] == 0 {
				que = append(que, ni)
			}
		}
	}
	if len(bytes) == len(ind) {
		return string(bytes)
	}

	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"})
	//alienOrder([]string{"z", "x"})
	//alienOrder([]string{"z", "x", "z"})
	//alienOrder([]string{"z", "z"})
	//alienOrder([]string{"ab", "adc"})
	//alienOrder([]string{"aac", "aabb", "aaba"})
	//alienOrder([]string{"ac", "ab", "zc", "zb"})
}
