package main

import "fmt"

func getPermutation(n int, k int) string {
	tmp := make([]byte, 0)
	for i := 1; i <= n; i++ {
		tmp = append(tmp, byte(i + '0'))
	}
	k--
	bytes := []byte{}
	for n > 0 {
		var f = factorial(n - 1)
		i := k/f
		bytes = append(bytes, tmp[i])
		copy(tmp[i:], tmp[i+1:])
		tmp = tmp[:n-1]
		n--
		k %= f
	}
	return string(bytes)
}

func factorial(x int) int {
	if x <= 1 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
}
