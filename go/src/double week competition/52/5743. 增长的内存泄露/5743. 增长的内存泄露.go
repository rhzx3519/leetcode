package main

import "fmt"

func memLeak(m1 int, m2 int) []int {
	//if m1 < m2 {
	//	return memLeak(m2, m1)
	//}
	//diff := m1 - m2
	for i := 1; ; i++ {
		if m1 >= m2 {
			if m1 < i {
				return []int{i, m1, m2}
			}
			m1 -= i
		} else {
			if m2 < i {
				return []int{i, m1, m2}
			}
			m2 -= i
		}
	}
}

func main() {
	//fmt.Println(memLeak(2, 2))
	//fmt.Println(memLeak(8, 11))
	fmt.Println(memLeak(5, 7))
}