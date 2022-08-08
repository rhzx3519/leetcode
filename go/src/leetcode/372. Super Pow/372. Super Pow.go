/**
@author ZhengHao Lou
Date    2021/12/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/super-pow/
思路：快速幂
a^112 = a^110 * a^2 = (a^11)^10*a^2
 */
const MOD = 1337
func superPow(a int, b []int) int {
	var ans = 1
	nb := len(b)
	for i := nb - 1; i >= 0; i-- {
		ans = ans*pow(a, b[i])%MOD
		a = pow(a, 10)
	}
	return ans
}

func pow(x, n int) int {
	var res = 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(superPow(2, []int{3}))
	fmt.Println(superPow(2, []int{1,0}))
	fmt.Println(superPow(1, []int{4,3,3,8,5,2}))
	fmt.Println(superPow(2147483647, []int{2,0,0}))
}