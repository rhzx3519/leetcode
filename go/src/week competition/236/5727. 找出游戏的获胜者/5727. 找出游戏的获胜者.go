package main

func findTheWinner(n int, k int) int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i+1
	}

	var cur = 0
	for len(a) != 1 {
		cur = (cur + k - 1) % len(a)
		copy(a[cur:], a[cur+1:])
		a = a[:len(a)-1]
	}
	//fmt.Println(a[0])
	return a[0]
}

func main() {
	findTheWinner(5, 2)
	findTheWinner(6, 5)
}
