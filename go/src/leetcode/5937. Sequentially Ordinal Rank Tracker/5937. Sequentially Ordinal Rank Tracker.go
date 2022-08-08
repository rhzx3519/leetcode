/**
@author ZhengHao Lou
Date    2021/12/13
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/sequentially-ordinal-rank-tracker/
 */
type Scene struct {
	name string
	score int
}

type SORTracker struct {
	i int
	scenes []Scene
}


func Constructor() SORTracker {
	return SORTracker{
		scenes: []Scene{},
	}
}


func (this *SORTracker) Add(name string, score int)  {
	scene := Scene{
		name: name,
		score: score,
	}
	i := upperBound(this.scenes, scene)
	this.scenes = append(this.scenes, Scene{})
	copy(this.scenes[i+1:], this.scenes[i:])
	this.scenes[i] = scene

	fmt.Println(this.scenes)
}


func (this *SORTracker) Get() string {
	defer func() {
		this.i++
	}()
	return this.scenes[this.i].name
}

func upperBound(scenes []Scene, x Scene) int {
	var m int
	l, r := 0, len(scenes)
	for l < r {
		m = l + (r - l)>>1
		if scenes[m].score < x.score || (scenes[m].score == x.score && scenes[m].name > x.name) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	tracker := Constructor()
	tracker.Add("bradford", 2)
	tracker.Add("branford", 3)
	tracker.Add("alps", 2)

}

/**
 * Your SORTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(name,score);
 * param_2 := obj.Get();
 */
