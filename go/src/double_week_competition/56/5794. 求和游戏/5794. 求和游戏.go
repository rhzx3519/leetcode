package main

import "fmt"

/**
思路：mapper[i]表示第i位为x时，num 中前一半数字的和 是否等于 后一半数字的和
题解：https://leetcode-cn.com/problems/sum-game/solution/alice-ke-jin-kai-gua-de-xie-beats-shuang-8gu8/
 */
func sumGame(num string) bool {
	var n = len(num)
	var c1, c2 int
	var s1, s2 int
	for i := range num {
		if i < n/2 {
			if num[i] == '?' {
				c1++
				continue
			}
			s1 += int(num[i] - '0')
		} else {
			if num[i] == '?' {
				c2++
				continue
			}
			s2 += int(num[i] - '0')
		}
	}

	if (c1 + c2)&1 == 1 { // case1：问号总数为奇数，则Alice肯定获胜
		return true
	}

	if s1 > s2 {
		if c1 >= c2 {	// 前半部？数>后半部？，有剩余问号的一侧只增加，则Alice赢
			return true
		}
		return (c2 - c1)/2*9 + s2 != s1
	} else if s1 < s2 {
		if c1 <= c2 {	// 前半部？数<后半部？，有剩余问号的一侧只增加，则Alice赢
			return true
		}
		return (c1 - c2)/2*9 + s1 != s2
	} else {
		return c1 != c2 	// s1==s2, 如果前半部？数==后半部？数，则Bob获胜
	}
}


func main() {
	fmt.Println(sumGame("5023"))
	fmt.Println(sumGame("25??"))
	fmt.Println(sumGame("?3295???"))
}
