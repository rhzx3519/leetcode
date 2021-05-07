package main

import (
	"fmt"
)

var (
	t1 = uint16(0b0111100000)
	t2 = uint16(0b0001111000)
	t3 = uint16(0b0000011110)
	c = []uint16{t1, t2, t3}
)


func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	var ans = 2*n

	var mapper = make(map[int]uint16)
	for _, pair := range reservedSeats {
		r, c := pair[0] - 1, pair[1] - 1
		if c > 0 {
			mapper[r] |= 1 << c
		}
	}

	for _, s := range mapper {
		ans -= 2 - count(s)
	}

	return ans
}

func count(seat uint16) int {
	cnt := 0

	for i := 0; i < len(c); i++ {
		if (seat & c[i]) == 0 {
			//fmt.Printf("%b\n%b\n\n", seat, c[i])
			cnt++
			i++
		}
	}

	return cnt
}

func main() {
	//var seat uint16
	//seat |= 1 << 1
	//seat |= 1 << 2
	//seat |= 1 << 7

	//seat |= 1 << 5

	//seat |= 1
	//seat |= 1 << 9
	//fmt.Println(count(seat))

	fmt.Println(maxNumberOfFamilies(3, [][]int{{1,2},{1,3},{1,8},{2,6},{3,1},{3,10}}))
	fmt.Println(maxNumberOfFamilies(2, [][]int{{2,1},{1,8},{2,6}}))
	fmt.Println(maxNumberOfFamilies(4, [][]int{{4,3},{1,4},{4,6},{1,7}}))
}
