package main

/*
*
https://leetcode.cn/problems/minimum-garden-perimeter-to-collect-enough-apples/?envType=daily-question&envId=2023-12-24

*/
/**
sn = 2n(n+1)(2n+1)
*/
func minimumPerimeter(neededApples int64) int64 {
	var n = int64(1)
	for 2*n*(n+1)*(2*n+1) < neededApples {
		n++
	}
	return n * 8
}
