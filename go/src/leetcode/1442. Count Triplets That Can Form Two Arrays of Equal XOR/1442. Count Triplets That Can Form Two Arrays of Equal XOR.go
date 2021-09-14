package main

func countTriplets(arr []int) int {
	n := len(arr)
	pre := make([]int, n+1)
	for i := 1; i <=n ;i++ {
		pre[i] = pre[i-1] ^ arr[i-1]
	}
	//fmt.Println(pre[2]^pre[0])
	var count int
	for j := 1; j < n; j++ {
		for i := j-1; i >= 0; i-- {
			var l = pre[j] ^ pre[i]
			for k := j; k < n; k++ {
				r := pre[k+1] ^ pre[j]
				if l == r {
					count++
				}
			}
		}
	}
	//fmt.Println("count:", count)
	return count
}

func main() {
	countTriplets([]int{2,3,1,6,7})
	countTriplets([]int{1,1,1,1,1})
	countTriplets([]int{2,3})
	countTriplets([]int{1,3,5,7,9})
	countTriplets([]int{7,11,12,9,5,2,7,17,22})
}