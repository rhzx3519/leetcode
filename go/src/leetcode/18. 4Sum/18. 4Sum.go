package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			l, r := j+1, n-1
			for l < r {
				val := nums[i] + nums[j] + nums[l] + nums[r]
				if val < target {
					l++
				} else if val > target {
					r--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				}
			}
			for j < n-1 && nums[j]==nums[j+1] {
				j++
			}
		}
		for i < n-1 && nums[i]==nums[i+1] {
			i++
		}
	}

	return ans
}

func kSum(nums []int, k, target int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	if n < k {
		return ans
	}
	if k > 2 {
		for i := 0; i < n; i++ {
			for _, tmp := range kSum(nums[i+1:], k-1, target-nums[i]) {
				if len(tmp) == k-1 {
					tmp = append(tmp, nums[i])
					ans = append(ans, tmp)
				}
			}

			for i < n-1 && nums[i]==nums[i+1] {
				i++
			}
		}
	}

	if k == 2 {
		l, r := 0, n-1
		for l < r {
			val := nums[l] + nums[r]
			if val > target {
				r--
			} else if val < target {
				l++
			} else {
				ans = append(ans, []int{nums[l], nums[r]})
				for l < r && nums[l]==nums[l+1] {
					l++
				}
				l++
			}
		}
	}

	return ans
}

func main() {
	//fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
	//fmt.Println(fourSum([]int{}, 0))

	fmt.Println(kSum([]int{1,0,-1,0,-2,2}, 4, 0))
	fmt.Println(kSum([]int{-2,-1,-1,1,1,2,2}, 4, 0))
}