/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/letter-tile-possibilities/
思路：暴力dfs
 */
const N = 26
func numTilePossibilities(tiles string) int {
	letters := make([]int, N)
	for i := range tiles {
		letters[int(tiles[i] - 'A')]++
	}
	fmt.Println(letters)
	var count int
	var dfs func([]int)
	dfs = func(letters []int) {
		for i := range letters {
			if letters[i] != 0 {
				count++
				letters[i]--
				dfs(letters)
				letters[i]++
			}
		}
	}
	dfs(letters)

	return count
}

func main() {
	fmt.Println(numTilePossibilities("AAB"))
	fmt.Println(numTilePossibilities("AAABBC"))
}
