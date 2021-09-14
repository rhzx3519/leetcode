package main

import "fmt"

const (
	BH = 0xF
	BM = 0x3F0
	N = 10

)

/**
k用来表示4位的hour，6位的minute
 */
func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	ans := []string{}

	var dfs func(i, k, on int)
	dfs = func(i, k, on int) {
		if on == 0 {
			var h, m = convert(k&BH), convert(k&BM >> 4)
			if h < 12 && m < 60 {
				str := fmt.Sprintf("%d:%02d", h, m)
				ans = append(ans, str)
			}
			return
		}

		for j := i; j < N; j++ {
			dfs(j+1, k|1<<j, on-1)
		}

	}
	dfs(0, 0, turnedOn)
	fmt.Println(ans)
	return ans
}

func convert(b int) int {
	var a, i int
	for b != 0 {
		if b&1 == 1 {
			a += 1<<i
		}
		i++
		b >>= 1
	}
	return a
}

// 方法2：预先生成所有的可能结果

var (
	mapper = make(map[int][]string)
)

func init() {
	for h := 0; h <	12; h++ {
		for m := 0; m < 60; m++ {
			var on = countOn(h) + countOn(m)
			if _, ok := mapper[on]; !ok {
				mapper[on] = []string{}
			}
			mapper[on] = append(mapper[on], fmt.Sprintf("%d:%02d", h, m))
		}
	}
}

func readBinaryWatch2(turnedOn int) []string {
	return mapper[turnedOn]
}

func countOn(x int) int {
	var i int
	for ; x > 0; x -= lowbit(x) {
		i++
	}
	return i
}

func lowbit(x int) int {
	return x & -x
}


func main() {
	fmt.Println(readBinaryWatch2(1))
	fmt.Println(readBinaryWatch2(9))
}
