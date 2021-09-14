package main

import (
	"fmt"
	"math"
)

var mod = int(math.Pow10(9) + 7)

func countPairs(deliciousness []int) int {
	mapper := make(map[int]int)
	for _, d := range deliciousness {
		mapper[d]++
	}
	ds := make([]int, 0)
	for k, _ := range mapper {
		ds = append(ds, k)
	}
	var ans int
	n := len(ds)
	for i := 0; i < n; i++ {
		if power(ds[i]) {
			ans = (ans + sum(mapper[ds[i]] - 1)) % mod
		}
		for j := i+1; j < n; j++ {
			s := ds[i] + ds[j]
			if power(s) {
				ans = (ans + mapper[ds[i]] * mapper[ds[j]]) % mod
			}
		}
	}
	return ans
}

func sum(n int) int {
	return (1+n)*n/2
}

func power(x int) bool {
	return x>0 && (x&(x-1)==0)
}

func main() {
	fmt.Println(countPairs([]int{1,3,5,7,9}))
	fmt.Println(countPairs([]int{1,1,1,3,3,3,7}))
	fmt.Println(countPairs([]int{2160,1936,3,29,27,5,2503,1593,2,0,16,0,3860,28908,6,2,15,49,6246,1946,23,105,7996,196,0,2,55,457,5,3,924,7268,16,48,4,0,12,116,2628,1468}))
}