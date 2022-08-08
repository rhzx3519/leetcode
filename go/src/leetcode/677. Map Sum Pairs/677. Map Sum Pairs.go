/**
@author ZhengHao Lou
Date    2021/11/14
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/map-sum-pairs/
 */
type MapSum struct {
	mapper map[string]int
}


func Constructor() MapSum {
	return MapSum{mapper: map[string]int{}}
}


func (this *MapSum) Insert(key string, val int)  {
	this.mapper[key] = val
}


func (this *MapSum) Sum(prefix string) int {
	var c int
	for k, v := range this.mapper {
		if strings.Index(k, prefix) == 0 {
			c += v
		}
	}
	return c
}

