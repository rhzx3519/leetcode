/**
@author ZhengHao Lou
Date    2021/11/15
*/
package main

type Dir [2]int

var (
	East = [2]int{1, 0}
	North = [2]int{0, 1}
	West = [2]int{-1, 0}
	South = [2]int{0, -1}
	dirs = [4]Dir{East, North, West, South}
	dirNames = [4]string{"East", "North", "West", "South"}
)

type Pos struct {
	x, y int
}

func (p Pos) position() []int {
	return []int{p.x, p.y}
}

type Robot struct {
	w, h int
	pos Pos
	dir int
}

func init() {

}

func Constructor(width int, height int) Robot {
	return Robot{
		w:   width,
		h:   height,
		pos: Pos{x: 0, y: 0},
		dir: 0,
	}
}


func (this *Robot) Move(num int)  {
	// prune
	cicle := this.w<<1 + this.h<<1 - 4
	num %= cicle
	if num == 0 { // corner
		if (this.pos.x == 0 && this.pos.y == 0) {
			this.dir = 3
		}
		if (this.pos.x == this.w-1 && this.pos.y == 0) {
			this.dir = 0
		}
		if (this.pos.x == this.w-1 && this.pos.y == this.h-1) {
			this.dir = 1
		}
		if (this.pos.x == 0 && this.pos.y == this.h-1) {
			this.dir = 2
		}
		return
	}

	for num != 0 {
		d := dirs[this.dir]
		nx := this.pos.x + d[0]*num
		ny := this.pos.y + d[1]*num
		if nx >= 0 && nx < this.w && ny >= 0 && ny < this.h {
			this.pos.x = nx
			this.pos.y = ny
			num = 0
			break
		}
		if nx < 0 {
			num -= this.pos.x
			this.pos.x = 0
		} else if nx >= this.w {
			num -= this.w - 1 - this.pos.x
			this.pos.x = this.w - 1
		} else if ny < 0 {
			num -= this.pos.y
			this.pos.y = 0
		} else if ny >= this.h {
			num -= this.h - 1 - this.pos.y
			this.pos.y = this.h - 1
		}
		this.dir = (this.dir + 1) % len(dirs)
	}
}


func (this *Robot) GetPos() []int {
	return this.pos.position()
}


func (this *Robot) GetDir() string {
	return dirNames[this.dir]
}


/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
