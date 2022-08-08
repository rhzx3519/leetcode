/**
@author ZhengHao Lou
Date    2021/11/28
*/
package main

import "fmt"

func getAverages(nums []int, k int) []int {
	n := len(nums)
	sums := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		sums[i] = nums[i-1] + sums[i-1]
	}
	var ans = make([]int, n)
	//var s int
	var t = k<<1 + 1
	for i := range nums {	// [i-k, i+k], i-k>=0, i+k <n, i+k -i+k + 1 = K<<1 + 1
		if i - k < 0 || i +k >= n {
			ans[i] = -1
			continue
		}

		ans[i] = (sums[i+k+1] - sums[i-k]) / t
	}
	return ans
}

func main() {
	fmt.Println(getAverages([]int{7,4,3,9,1,8,5,2,6}, 3))
	fmt.Println(getAverages([]int{100000}, 0))
	fmt.Println(getAverages([]int{8}, 100000))
}
