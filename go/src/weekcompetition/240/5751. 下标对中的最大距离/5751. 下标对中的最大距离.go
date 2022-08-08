package main

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	ans := 0
	for i := 0; i < len(nums1); i++ {
		if i+1 >= len(nums2) {
			break
		}
		t := UpperBound(nums2[i+1:], nums1[i])
		if t > ans {
			ans = t
		}
	}
	return ans
}

/**
返回第一个arr中第一个大于的元素下标，
如果k大于arr中的所有元素，则返回arr的长度
*/
func UpperBound(arr []int, k int) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if arr[mid] >= k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	//fmt.Println(maxDistance([]int{9914,9434,8808,8041,7548,6727,5637,4635,2937,607,384,335},
	//[]int{9980,9826,9620,9600,9455,9448,9374,9367,9278,9251,9083,8987,8952,8932,8762,8705,8595,8460}))
	fmt.Println(UpperBound([]int{5,3,3,2,2,1}, 1))
	fmt.Println(maxDistance([]int{55,30,5,4,2}, []int{100,20,10,10,5}))
	fmt.Println(maxDistance([]int{30,29,19,5}, []int{25,25,25,25,25}))
	fmt.Println(maxDistance([]int{5,4}, []int{3,2}))
	fmt.Println(maxDistance([]int{9819,9508,7398,7347,6337,5756,5493,5446,5123,3215,1597,774,368,313}, []int{9933,9813,9770,9697,9514,9490,9441,9439,8939,8754,8665,8560}))
}
