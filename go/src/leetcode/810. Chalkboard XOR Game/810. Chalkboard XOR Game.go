package main

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/chalkboard-xor-game/solution/hei-ban-yi-huo-you-xi-by-leetcode-soluti-eb0c/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。