/**
@author ZhengHao Lou
Date    2021/10/02
*/
package main

import (
	"fmt"
)

func numOfPairs(nums []string, target string) int {
	var count int
	mapper := map[string]int{}
	for _, num := range nums {
		if len(num) >= len(target) {
			continue
		}

		//t := target[len(target) - len(num):]
		//fmt.Println(t)
		if num == target[:len(num)] {
			count += mapper[target[len(num):]]
		}
		if num == target[len(target) - len(num):] {
			count += mapper[target[:len(target) - len(num)]]
		}
		mapper[num]++
	}
	return count
}

func main() {
	fmt.Println(numOfPairs([]string{"777","7","77","77"}, "7777"))
	fmt.Println(numOfPairs([]string{"123","4","12","34"}, "1234"))
	fmt.Println(numOfPairs([]string{"1","1","1"}, "11"))
}
