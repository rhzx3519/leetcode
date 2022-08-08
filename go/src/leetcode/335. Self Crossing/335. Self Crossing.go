/**
@author ZhengHao Lou
Date    2021/10/29
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/self-crossing/
作者：AC_OIer
链接：https://leetcode-cn.com/problems/self-crossing/solution/gong-shui-san-xie-fen-qing-kuang-tao-lun-zdrb/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func isSelfCrossing(d []int) bool {
	n := len(d)
	if n < 4 {
		return false;
	}

	for i := 3; i < n; i++ {
		if d[i] >= d[i - 2] && d[i - 1] <= d[i - 3] {
			return true
		}

		if i >= 4 && d[i - 1] == d[i - 3] && d[i] + d[i - 4] >= d[i - 2] {
			return true
		}
		if i >= 5 && d[i - 1] <= d[i - 3] && d[i - 2] > d[i - 4] && d[i] + d[i - 4] >= d[i - 2] && d[i - 1] + d[i - 5] >= d[i - 3] {
			return true
		}
	}
	strings.Join()
	fmt.Printf()
	return false;
}

func main() {
	fmt.Println(isSelfCrossing([]int{2,1,1,2}))
	fmt.Println(isSelfCrossing([]int{1,2,3,4}))
	fmt.Println(isSelfCrossing([]int{1,1,1,1}))

	fmt.Println(isSelfCrossing([]int{1,1,2,1,1}))
}
