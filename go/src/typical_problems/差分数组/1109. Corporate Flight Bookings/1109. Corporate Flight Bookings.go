package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, b := range bookings {
		i, j, s := b[0]-1, b[1]-1, b[2]
		diff[i] += s
		diff[j+1] -= s
	}
	nums := make([]int, n)
	nums[0] = diff[0]
	for i := 1; i < n; i++ {
		nums[i] = diff[i] + nums[i-1]
	}
	fmt.Println(diff)
	fmt.Println(nums)
	return nums
}

func main() {
	corpFlightBookings([][]int{{1,2,10},{2,3,20},{2,5,25}}, 5)
	corpFlightBookings([][]int{{1,2,10},{2,2,15}}, 2)
}
