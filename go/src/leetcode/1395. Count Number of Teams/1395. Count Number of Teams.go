package main

import "fmt"

func numTeams(rating []int) int {
	n := len(rating)
	st := []int{}
	asc := make([]int, n)
	desc := make([]int, n)
	for i := 0; i < n; i++ {
		asc[i] = -1
		desc[i] = -1
	}

	for i := range rating {
		for len(st) > 0 && rating[st[len(st)-1]] < rating[i] {
			asc[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	fmt.Println(asc)

	for i := n-1; i >= 0; i-- {
		for len(st) > 0 && rating[st[len(st)-1]] > rating[i] {
			desc[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	fmt.Println(desc)

	var count int

	// dp[i]表示以rating[i]结尾的递增序列的个数(递增序列的元素个数大于等于2)
	// asc[i] != -1
	for i := 0; i < n; i++ {
		if asc[i] != -1 && asc[asc[i]] != -1 {
			count++
		}
	}
	for i := 0; i >= 0; i-- {
		if desc[i] != -1 && desc[desc[i]] != -1 {
			count++
		}
	}


	return count
}

func main() {
	numTeams([]int{2,5,3,4,1})
}
