package main

import "fmt"

/*
*
https://leetcode.cn/problems/sum-of-distances/
0,2,3
s = 5
s - 2*2 + 1*2 = 3
s - 1*1 + 2*1 = 4
*/
func distance(nums []int) []int64 {
	n := len(nums)
	arr := make([]int64, n)
	cnt := map[int][]int{}
	for i, x := range nums {
		cnt[x] = append(cnt[x], i)

	}

	for _, d := range cnt {
		m := len(d)
		var s int64
		for i := 1; i < m; i++ {
			s += int64(d[i] - d[0])
		}
		arr[d[0]] = s
		for i := 1; i < m; i++ {
			t := int64(d[i] - d[i-1])
			s = s - int64(m-i)*t + int64(i)*t
			arr[d[i]] = s
		}
	}

	return arr
}

func main() {
	fmt.Println(distance([]int{1, 3, 1, 1, 2}))
	fmt.Println(distance([]int{0, 5, 3}))
}
