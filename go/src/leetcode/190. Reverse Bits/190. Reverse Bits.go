/**
思路：
1. 取第i位
	x = (n>>i)&1
2. 设置第i位为x
	n ^= (n&(1<<i) ^ (x<<i)
 */

package main

import "fmt"

func reverseBits(num uint32) uint32 {
	for i := 0; i < 16; i++ {
		x := (num>>(31-i)) & 1
		y := (num>>i) & 1
		num ^= (num&(1<<(31-i))) ^ (y<<(31-i))
		num ^= (num&(1<<i)) ^ (x<<i)
	}
	return num
}

func main() {
	fmt.Printf("%b", reverseBits(uint32(00000010100101000001111010011100)))
}