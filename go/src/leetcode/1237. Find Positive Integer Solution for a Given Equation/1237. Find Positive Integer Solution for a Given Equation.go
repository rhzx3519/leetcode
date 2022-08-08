/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

/**
https://leetcode-cn.com/problems/find-positive-integer-solution-for-a-given-equation/
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */
const N = 1000
func findSolution(customFunction func(int, int) int, z int) [][]int {
	var ans = [][]int{}
	for x := 1; x <= N; x++ {
		for y := 1; y <= N; y++ {
			if customFunction(x, y) == z {
				ans = append(ans, []int{x,y})
			}
		}
	}
	return ans
}
