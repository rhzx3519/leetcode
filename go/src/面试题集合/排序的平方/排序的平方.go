package main

import "fmt"

/**
题目：给一个有序数组比如[-4,-1,0,6,10] 输出排序的平方[0,1,16,36,100]，时间复杂度为N的做法
 */
func sortedArraySquare(arr []int) []int {
	n := len(arr)
	var i int
	for i < n && arr[i] < 0 {
		i++
	}
	ans := []int{}
	j, k := i-1, i
	for j >= 0 || k < n {
		if j >= 0 && k < n {
			if -arr[j] < arr[k] {
				ans = append(ans, arr[j]*arr[j])
				j--
			} else {
				ans = append(ans, arr[k]*arr[k])
				k++
			}
		} else if j >= 0 {
			ans = append(ans, arr[j]*arr[j])
			j--
		} else {
			ans = append(ans, arr[k]*arr[k])
			k++
		}
	}

	return ans
}

func main() {
	// [0,1,16,36,100]
	fmt.Println(sortedArraySquare([]int{-4,-1,0,6,10}))
}

