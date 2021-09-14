package main

import "fmt"

/**
https://leetcode-cn.com/problems/circular-array-loop/submissions/
思路: 快慢指针
 */
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	var next func(i int) int
	next = func(i int) int {
		return ((i + nums[i])%n + n)%n
	}


	for i := range nums {
		slow, fast := i, next(i)
		for {
			if nums[slow]*nums[fast] < 0 || nums[fast]*nums[next(fast)] < 0 {
				break
			}
			if slow == fast {
				if slow == next(slow) {
					break
				} else {
					return true
				}
			}

			slow = next(slow)
			fast = next(next(fast))
		}
	}
	return false
}

func main() {
	fmt.Println(circularArrayLoop([]int{2,-1,1,2,2}))
	fmt.Println(circularArrayLoop([]int{-1,2}))
	fmt.Println(circularArrayLoop([]int{-2,1,-1,-2,-2}))
}
