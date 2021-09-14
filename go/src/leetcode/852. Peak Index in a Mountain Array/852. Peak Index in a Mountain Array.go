package main

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	l, r := 1, n-2
	var mid, ans int
	for l <= r {
		mid = l + (r-l)>>1
		if arr[mid] > arr[mid+1] {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}
