package main

/*
*
https://leetcode.cn/problems/adding-two-negabinary-numbers/description/
思路：模拟
-(-2)^i = (-2)^i + (-2)^(i+1)
*/
func addNegabinary(arr1 []int, arr2 []int) []int {
	var ans []int
	i, j := len(arr1)-1, len(arr2)-1
	for c := 0; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
		x := c
		c = 0

		if i >= 0 {
			x += arr1[i]
		}
		if j >= 0 {
			x += arr2[j]
		}

		if x >= 2 {
			x -= 2
			c += -1 // 向高位进-1
		} else if x == -1 {
			x = 1
			c += 1 // 向高位进1
		}
		ans = append(ans, x)
	}

	for i := len(ans) - 1; i > 0 && ans[i] == 0; i-- {
		ans = ans[:len(ans)-1]
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func main() {
	addNegabinary([]int{1}, []int{1, 1, 0, 1})
}
