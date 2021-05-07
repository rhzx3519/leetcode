/**
https://leetcode-cn.com/problems/single-number-ii/
思路：将数字分解成32位的比特串，每一位出现的次数分别是出现3次的数字+出现1次的数字
所以b = s(b)%3
 */
package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var t int32
	var b int32
	for i := 0; i < 32; i++ {
		b = int32(0)
		for _, num := range nums {
			//fmt.Printf("%d, %b\n", num, num)
			fmt.Println((num>>i)&1)
			b += (int32(num)>>i)&1
		}
		if b%3 > 0 {
			t |= 1 << i
		}
	}
	return int(t)
}

func main() {
	fmt.Println(singleNumber([]int{-2,-2,1,1,4,1,4,4,-4,-2}))
	//fmt.Println(singleNumber([]int{0,1,0,1,0,1,99}))

}