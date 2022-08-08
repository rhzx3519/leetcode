/**
@author ZhengHao Lou
Date    2022/07/13
*/
package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	var i int
	for i < len(asteroids) {
		if i-1 < 0 {
			i++
			continue
		}
		a, b := asteroids[i-1], asteroids[i]
		if a*b > 0 { // ll, rr
			i++
		} else if a < 0 { // lr
			i++
		} else { // rl
			if abs(a) == abs(b) {
				copy(asteroids[i-1:], asteroids[i+1:])
				asteroids = asteroids[:len(asteroids)-2]
				i--
			} else if abs(a) > abs(b) {
				copy(asteroids[i:], asteroids[i+1:])
				asteroids = asteroids[:len(asteroids)-1]
			} else {
				copy(asteroids[i-1:], asteroids[i:])
				asteroids = asteroids[:len(asteroids)-1]
				i--
			}
		}
	}

	fmt.Println(asteroids)
	return asteroids
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	asteroidCollision([]int{5, 10, -5})
	asteroidCollision([]int{8, -8})
	asteroidCollision([]int{10, 2, -5})
	asteroidCollision([]int{-2, -1, 1, 2})
	asteroidCollision([]int{-2, 2, -1, -2}) // [-2]
	asteroidCollision([]int{1, 1, -1, -2})  // [-2]
}
