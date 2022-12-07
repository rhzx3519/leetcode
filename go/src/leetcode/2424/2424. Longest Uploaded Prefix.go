/**
@author ZhengHao Lou
Date    2022/10/04
*/
package main

type LUPrefix struct {
	index  int
	videos []bool
}

func Constructor(n int) LUPrefix {
	return LUPrefix{videos: make([]bool, n)}
}

func (this *LUPrefix) Upload(video int) {
	this.videos[video] = true
	j := this.index
	for j < len(this.videos) && this.videos[j] {
		j++
	}
	this.index = j
}

func (this *LUPrefix) Longest() int {
	return this.index
}
