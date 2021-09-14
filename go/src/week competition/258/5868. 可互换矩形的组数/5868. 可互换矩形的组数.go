package main

import "fmt"

func interchangeableRectangles(rectangles [][]int) int64 {
	mapper := make(map[float64]int64)
	mapper[float64(0)] = int64(0)
	var count int64

	for _, rect := range rectangles {
		wh := float64(rect[0]) / float64(rect[1])
		if c, ok := mapper[wh]; ok {
			count += c
		}
		mapper[wh]++
	}

	return count
}

func main() {
	fmt.Println(interchangeableRectangles([][]int{{4,8},{3,6},{10,20},{15,30}}))
	fmt.Println(interchangeableRectangles([][]int{{4,5},{7,8}}))
}
