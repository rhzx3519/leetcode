package main

import "fmt"

func getHint(secret string, guess string) string {
	var x, y int
	n := len(secret)
	m1 := make(map[byte][]int)
	m2 := make(map[byte][]int)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			x++
		} else {
			m1[secret[i]] = append(m1[secret[i]], i)
			m2[guess[i]] = append(m2[guess[i]], i)
		}
	}

	for d, p2 := range m2 {
		if p1, ok := m1[d]; ok {
			y += min(len(p1), len(p2))
		}
	}

	return fmt.Sprintf("%dA%dB", x, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
}