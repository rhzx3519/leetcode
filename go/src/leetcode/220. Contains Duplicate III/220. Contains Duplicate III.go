package main

import (
	"fmt"
	"math"
)

/**
思路：大小为t+1的桶, x-t <= x <= x+t
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	w := int64(t)
	mapper := make(map[int64]int)

	bucketId := func(x int) int64 {
		return int64(math.Floor(float64(x) / float64(w+1)))
	}

	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}

	for i, x := range nums {
		id := bucketId(x)
		if _, ok := mapper[id]; ok {
			return true
		}
		if left, ok := mapper[id-1]; ok && abs(x - left) <= t {
			return true
		}
		if right, ok := mapper[id+1]; ok && abs(x - right) <= t {
			return true
		}
		mapper[id] = x
		if i >= k {
			delete(mapper, bucketId(nums[i-k]))
		}
	}
	return false
}


func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,2,3,1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,0,1,1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{2147483647,-1,2147483647}, 1, 2147483647))
}