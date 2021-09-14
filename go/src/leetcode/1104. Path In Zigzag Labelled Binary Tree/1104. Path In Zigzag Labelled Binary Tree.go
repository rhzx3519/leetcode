package main

import (
	"fmt"
	"math"
)

func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	var lv = 0
	for pow2(lv) <= label {
		lv++
	}
	fmt.Println(lv)
	var parent int
	var path = []int{}
	for i := lv-1; i > 0; i-- {
		path = append(path, label)
		if label&1 == 0 { // even
			parent = label >> 1
		} else {
			parent = (label - 1) >> 1
		}

		l, r := pow2(i-1), pow2(i) - 1
		parent = mirror(l, r, parent)

		label = parent
	}
	path = append(path, 1)
	reverse(path)

	return path
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func mirror(l, r, x int) int {
	var m = l + (r-l)>>1
	return m<<1 - x + 1
}

func pow2(n int) int {
	return int(math.Pow(float64(2), float64(n)))
}

func main() {
	fmt.Println(pathInZigZagTree(14))
	fmt.Println(pathInZigZagTree(26))
	fmt.Println(pathInZigZagTree(3))
}
