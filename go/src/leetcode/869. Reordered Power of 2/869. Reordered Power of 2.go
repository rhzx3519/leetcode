/**
@author ZhengHao Lou
Date    2021/10/28
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/reordered-power-of-2/
 */
func reorderedPowerOf2(n int) bool {
	var digits = []int{}
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	N := len(digits)
	vis := make([]bool, N)
	var backtrace func(idx int, nums []int) bool
	backtrace = func(idx int, nums []int) bool {
		if len(nums) == N {
			return check(nums)
		}

		for i := 0; i < N; i++ {
			if !vis[i] {
				if digits[i] == 0 && len(nums) == 0 {
					continue
				}
				vis[i] = true
				nums = append(nums, digits[i])
				if backtrace(i, nums) {
					return true
				}
				nums = nums[:len(nums) - 1]
				vis[i] = false
			}
		}
		return false
	}

	return backtrace(0, []int{})
}


func check(nums []int) bool {
	var x int
	for _, num := range nums {
		x = x*10 + num
	}
	return x&(x-1) == 0
}

func reorderedPowerOf2_2(n int) bool {
	nums := []int{}
	for i := 0; i <= 30; i++ {
		nums = append(nums, 1<<i)
	}
	mapper := make([]int, 10)
	for n != 0 {
		mapper[n%10]++
		n /= 10
	}

	for _, n2 := range nums {
		mp := make([]int, 10)
		for n2 != 0 {
			mp[n2 % 10]	++
			n2 /= 10
		}
		var f = true
		for i := 0; i < 10; i++ {
			if mapper[i] != mp[i] {
				f = false
				break
			}
		}
		if f {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(reorderedPowerOf2_2(1))
	fmt.Println(reorderedPowerOf2_2(10))
	fmt.Println(reorderedPowerOf2_2(16))
	fmt.Println(reorderedPowerOf2(24))
	fmt.Println(reorderedPowerOf2(46))
}
