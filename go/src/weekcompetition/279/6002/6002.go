/**
@author ZhengHao Lou
Date    2022/02/06
*/
package main

const UNIT int = 64

type Bitset struct {
	size  int
	bits  []uint64
	count int
}

func Constructor(size int) Bitset {
	l := size/UNIT + 1
	return Bitset{bits: make([]uint64, l), size: size}
}

func (this *Bitset) Fix(idx int) {
	i, r := idx/UNIT, idx%UNIT
	if this.bits[i]>>r&1 == 0 {
		this.count++
	}
	this.bits[i] |= 1 << r

}

func (this *Bitset) Unfix(idx int) {
	i, r := idx/UNIT, idx%UNIT
	if this.bits[i]>>r&1 == 1 {
		this.count--
	}
	this.bits[i] &= ^(1 << r)
}

func (this *Bitset) Flip() {
	for i := range this.bits {
		this.bits[i] = ^this.bits[i]
	}
	this.count = this.size - this.count
}

func (this *Bitset) All() bool {
	for i := 0; i < this.size/UNIT; i++ {
		if this.bits[i] != 1<<64-1 {
			return false
		}
	}
	i, r := this.size/UNIT, this.size%UNIT
	for k := 0; k < r; k++ {
		if this.bits[i]>>k&1 == 0 {
			return false
		}
	}

	return true
}

func (this *Bitset) One() bool {
	return this.Count() > 0
}

func (this *Bitset) Count() int {
	return this.count
}

func (this *Bitset) ToString() string {
	var bytes []byte
	for i := 0; i < this.size/UNIT; i++ {
		for j := 0; j < UNIT; j++ {
			if this.bits[i]>>j&1 == 0 {
				bytes = append(bytes, '0')
			} else {
				bytes = append(bytes, '1')
			}
		}
	}
	i, r := this.size/UNIT, this.size%UNIT
	for k := 0; k < r; k++ {
		if this.bits[i]>>k&1 == 0 {
			bytes = append(bytes, '0')
		} else {
			bytes = append(bytes, '1')
		}
	}
	return string(bytes)
}

func main() {
	obj := Constructor(5)
	//obj.Flip()
	//obj.Unfix(1)
	//obj.Fix(1)
	//obj.Fix(1)
	//obj.Unfix(1)

	obj.Fix(3)
	obj.Fix(1)
	obj.Flip()
	obj.All()
	obj.Unfix(0)
	obj.Flip()
	obj.One()
	obj.Unfix(0)
	obj.Count()
	obj.ToString()
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
