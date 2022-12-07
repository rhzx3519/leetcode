/**
@author ZhengHao Lou
Date    2022/10/05
*/
package main

/**
https://leetcode.cn/problems/minimize-xor/
*/
func minimizeXor(num1 int, num2 int) int {
	n1, n2 := c1(num1), c1(num2)
	if n1 == n2 {
		return num1
	} else if n1 > n2 {
		for k := 0; k < n1-n2; k++ {
			num1 = num1 & (num1 - 1)
		}
		return num1
	}
	var i, t = 0, num1
	for k := 0; k < n2-n1; k++ {
		for ; (num1>>i)&1 == 1; i++ {
		}
		t |= 1 << i
		i++
	}
	return t
}

func c1(n int) int {
	var b int
	for n != 0 {
		n = n & (n - 1)
		b++
	}
	return b
}

func main() {
	//minimizeXor(1, 12)
	//minimizeXor(25, 72) // 24
	//minimizeXor(65, 84) // 67
	minimizeXor(8, 75) // 15
}
