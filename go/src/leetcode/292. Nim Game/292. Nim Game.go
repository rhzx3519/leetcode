package main

/**
思路：考虑最小的问题集合，先手，
1. 剩余5个，先手拿1个必胜
2. 剩余4个，无论怎么拿，都必输
3。剩余1、2、3个，必胜
所以只要n%4 != 0, 都是必胜的。
 */
func canWinNim(n int) bool {
	return n%4 != 0
}
