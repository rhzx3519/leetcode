package main

import "fmt"

func openLock(deadends []string, target string) int {
	start := "0000"
	dead := make(map[string]bool)
	for _, s := range deadends {
		dead[s] = true
	}
	if _, ok := dead[start]; ok {
		return -1
	}
	dead[start] = true

	var step int
	que := []string{start}
	for len(que) > 0 {
		for sz := len(que); sz > 0; sz-- {
			s := que[0]
			que = que[1:]

			if s == target {
				return step
			}

			for i := range s {
				d := int(s[i] - '0')
				for _, j := range []int{-1, 1} {
					nd := mod(d + j, 10)
					bytes := []byte(s)
					bytes[i] = byte('0' + nd)
					t := string(bytes)
					if _, ok := dead[t]; ok {
						bytes[i] = s[i]
						continue
					}
					dead[t] = true
					que = append(que, t)
				}
			}
		}
		step++
	}

	return -1
}

func mod(x, m int) int {
	if x >= 0 {
		return x % m
	}
	for ; x < 0; x += m {}
	return x % m
}

func main() {
	fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
	fmt.Println(mod(10, 10))
}