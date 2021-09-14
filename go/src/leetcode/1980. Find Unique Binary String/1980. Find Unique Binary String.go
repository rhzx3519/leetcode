package main

import "fmt"

func findDifferentBinaryString(nums []string) string {
	var x int
	var mapper = make(map[int]bool)
	for _, num := range nums {
		x = bin(num)
		mapper[x] = true
	}
	n := len(nums)
	for i := 0; i <= 1<<n - 1; i++ {
		if _, ok := mapper[i]; !ok {
			return printf(i, n)
		}
	}
	//fmt.Printf("%016b\n", x)
	return ""
}

func printf(x, n int) string {
	bytes := []byte{}
	for i := 0; i < n; i++ {
		bytes = append(bytes, byte(x&1 + '0'))
		x >>= 1
	}
	l, r := 0, n-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func bin(num string) int {
	var x int
	n := len(num)
	for i := range num {
		d := int(num[i] - '0')
		if d == 1 {
			x |= 1 << (n-1-i)
		}
	}
	return x
}

func main() {
	//findDifferentBinaryString([]string{"01", "10"})
	//findDifferentBinaryString([]string{"00", "01"})
	findDifferentBinaryString([]string{"001","000","100"})
}
