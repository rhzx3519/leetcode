/**
@author ZhengHao Lou
Date    2022/01/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-good-people-based-on-statements/
思路：状态枚举
枚举所有的状态1<<n-1种，如果第i个人是好人，则状态应该和statements[i]一直；否则矛盾
*/
func maximumGood(statements [][]int) int {
	var ans int
	n := len(statements)
	mask := 1<<n - 1
next:
	for subset := mask; subset != 0; subset = (subset - 1) & mask {
		var good int
		for i := 0; i < n; i++ {
			if subset>>i&1 == 1 { // i is good
				for j, st := range statements[i] {
					if st < 2 && st != subset>>j&1 {
						continue next
					}
				}
				good++
			}
		}
		fmt.Println(subset)
		if good > ans {
			ans = good
		}
	}
	//fmt.Println(ans)
	return ans
}

func main() {
	maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}})
}
