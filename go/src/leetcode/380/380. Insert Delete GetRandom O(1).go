/**
@author ZhengHao Lou
Date    2022/04/13
*/
package main

import "math/rand"

/**
https://leetcode-cn.com/problems/insert-delete-getrandom-o1/
*/
type RandomizedSet struct {
    nums []int
    mp   map[int]int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        mp: map[int]int{},
    }
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.mp[val]; ok {
        return false
    }
    this.mp[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.mp[val]; !ok {
        return false
    }
    i := this.mp[val]
    delete(this.mp, val)

    if i != len(this.nums)-1 {
        this.nums[i], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[i]
        this.mp[this.nums[i]] = i
    }

    this.nums = this.nums[:len(this.nums)-1]
    return true
}

func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
