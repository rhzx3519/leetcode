package main

import "fmt"

func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make([]int, len(candiesCount)+1)
	for i := 1; i <= len(candiesCount); i++ {
		sum[i] = sum[i-1] + candiesCount[i-1]
	}

	ans := []bool{}
	for _, q := range queries {
		i, day, cap := q[0], q[1]+1, q[2]
		l := sum[i]/cap + 1		// 吃到第i类糖最短时间
		r := sum[i+1] 			// 吃到第i类糖最长时间
		ans = append(ans, day >=l && day <= r)
	}
	fmt.Println(ans)
	return ans
}



func main() {
	canEat([]int{7,4,5,3,8}, [][]int{{0,2,2},{4,2,4},{2,13,1000000000}})
}