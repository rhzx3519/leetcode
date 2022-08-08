/**
@author ZhengHao Lou
Date    2022/02/19
*/
package main

import "fmt"

/**
28: 2, 4, 6, 8, 10, 12
*/
func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum&1 != 0 {
		return []int64{}
	}
	coins := []int64{}
	for s := int64(0); s < finalSum; {
		s = s + 2
		coins = append(coins, s)
	}

	f := make(map[int64]int64)
	counter := make(map[int64][]int64)

	for _, c := range coins {
		for i := finalSum; i >= c; i -= 2 { // 遍历顺序不同
			if f[i-c]+1 >= f[i] {
				f[i] = f[i-c] + 1
				counter[i] = append([]int64{c}, counter[i-c]...)
			}
		}
	}

	fmt.Println(counter)
	fmt.Println(f)
	return []int64{}
}

func main() {
	maximumEvenSplit(71252)
}
