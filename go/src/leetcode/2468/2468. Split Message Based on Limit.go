/**
@author ZhengHao Lou
Date    2022/11/15
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/split-message-based-on-limit/
*/
func splitMessage(message string, limit int) []string {
	n := len(message)
	var k, cap, tailLen int
	for k = 1; ; k++ {
		if k < 10 {
			tailLen = 5
		} else if k < 100 {
			if k == 10 {
				cap -= 9 // 前面9个块的长度减1
			}
			tailLen = 7
		} else if k < 1000 {
			if k == 100 {
				cap -= 99 // 前面99个块的长度减1
			}
			tailLen = 9
		} else {
			if k == 1000 {
				cap -= 999
			}
			tailLen = 11
		}
		if tailLen >= limit {
			return nil
		}
		cap += limit - tailLen
		if cap >= n {
			break
		}
	}
	var ans []string
	for i := 1; i <= k; i++ {
		tail := fmt.Sprintf("<%v/%v>", i, k)
		if i == k {
			ans = append(ans, message+tail)
		} else {
			m := limit - len(tail)
			ans = append(ans, message[:m]+tail)
			message = message[m:]
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	splitMessage("this is really a very awesome message", 9)
}
