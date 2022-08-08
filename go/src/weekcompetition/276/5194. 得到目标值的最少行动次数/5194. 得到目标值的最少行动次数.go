/**
@author ZhengHao Lou
Date    2022/01/16
*/
package main

func minMoves(target int, maxDoubles int) int {
	var op int
	for target != 1 && maxDoubles != 0 {
		op += target%2 + 1
		target /= 2
		maxDoubles--
	}
	op += target - 1

	return op
}

func main() {
	//minMoves(5, 0)
	minMoves(19, 2)
	//fmt.Println(minMoves(10, 4))
	//fmt.Println(minMoves(766972377, 92))
}
