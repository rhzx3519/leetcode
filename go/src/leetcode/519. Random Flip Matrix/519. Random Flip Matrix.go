/**
@author ZhengHao Lou
Date    2021/11/27
*/
package main

import (
	"math/rand"
)

/**
https://leetcode-cn.com/problems/random-flip-matrix/
 */
type Solution struct {
	m, n int
	z int
	mp map[int]int
}


func Constructor(m int, n int) Solution {
	return Solution{
		m: m,
		n: n,
		z: m*n,
		mp: make(map[int]int),
	}
}


/**
z = m*n, 使用arr[0:z]来记录等于0的位置，r = i/n, c = i%n
每次Flip, 随机生成[0, z)之间的数 i, 则返回的等于0的下标位置为 r = arr[i]/n, c = arr[i]%n,
将置为1的下标放到arr[z-1], 保证arr[0:z]始终为0的位置下标
Flip完之后，z的数量减一,
实现时，使用哈希表记录被Flip的位置
 */
func (this *Solution) Flip() []int {
	i := rand.Intn(this.z)
	k := i
	if tmp, ok := this.mp[i]; ok {
		k = tmp
	}
	swap := this.z - 1
	if tmp, ok := this.mp[this.z - 1]; ok {
		swap = tmp
	}

	this.mp[i] = swap
	this.mp[this.z - 1] = k
	this.z--

	return []int{k/this.n, k%this.n}
}


func (this *Solution) Reset()  {
	this.mp = make(map[int]int)
	this.z = this.m*this.n
}

func main() {
	obj := Constructor(3,1)
	obj.Flip()
	obj.Flip()
	obj.Flip()
	obj.Reset()
	obj.Flip()
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
