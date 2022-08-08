/**
@author ZhengHao Lou
Date    2021/12/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/stickers-to-spell-word/
思路：完全背包，状态压缩

 */
func minStickers(stickers []string, target string) int {
	n := len(target)
	f := make([]int, 1<<n)	// f[i]表示当前状态为i时，最少的贴纸数量
	for i := range f {
		f[i] = -1
	}
	f[0] = 0

	for _, sticker := range stickers {
		for status := 0; status < (1<<n); status++ {
			if f[status] == -1 {
				continue
			}
			curStatus := status
			for i := range sticker {
				for j := 0; j < n; j++ {
					if sticker[i] == target[j] && curStatus & (1<<j) == 0 {
						curStatus |= 1<<j
						break
					}
				}
			}
			if f[curStatus] == -1 {
				f[curStatus] = f[status] + 1
			} else {
				f[curStatus] = min(f[curStatus], f[status] + 1)
			}
		}
	}

	return f[1<<n - 1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minStickers([]string{"with", "example", "science"}, "thehat"))
	fmt.Println(minStickers([]string{"notice", "possible"}, "basicbasic"))
}