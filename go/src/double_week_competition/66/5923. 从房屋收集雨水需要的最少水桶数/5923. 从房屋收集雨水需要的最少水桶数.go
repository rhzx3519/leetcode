/**
@author ZhengHao Lou
Date    2021/11/27
*/
package main

import "fmt"

func minimumBuckets(street string) int {
	arr := []int{}
	for i, c := range street {
		if byte(c) == 'H' {
			arr = append(arr, i)
		}
	}
	n := len(arr)
	var last = -1
	var bucket int
	var i int
	for i < n {
		//fmt.Println(last)
		if arr[i] + 1 < len(street) && street[arr[i] + 1] != 'H' {	// 优先放置在right
			last = arr[i] + 1
			bucket++
		} else if arr[i] - 1 >= 0 && street[arr[i] - 1] != 'H' {
			last = arr[i] - 1
			bucket++
		} else {
			return -1
		}
		i++
		if i < n && arr[i] == last + 1 {
			i++
		}
	}

	return bucket
}

func main() {
	fmt.Println(minimumBuckets("H..H"))			// 2
	fmt.Println(minimumBuckets(".H.H."))			// 1
	fmt.Println(minimumBuckets(".HHH."))			// -1
	fmt.Println(minimumBuckets("H"))				// -1
	fmt.Println(minimumBuckets("."))				// 0
	fmt.Println(minimumBuckets(".HH.H.H.H..")) 	// 3

}

