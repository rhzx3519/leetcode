/**
@author ZhengHao Lou
Date    2022/08/20
*/
package main

const N = 26

func shiftingLetters(s string, shifts [][]int) string {
	bytes := []byte(s)
	n := len(s)
	diff := make([]int, n+1)
	for _, shift := range shifts {
		l, r := shift[0], shift[1]
		if shift[2] == 0 {
			diff[l]--
			diff[r+1]++
		} else {
			diff[l]++
			diff[r+1]--
		}
	}

	sf := make([]int, n)
	sf[0] = (diff[0]%N + N) % N
	for i := 1; i < n; i++ {
		sf[i] = diff[i] + sf[i-1]
		sf[i] = (sf[i]%N + N) % N
	}

	//fmt.Println(sf)
	for i := range s {
		bytes[i] = byte('a' + (int(bytes[i]-'a')+sf[i]+N)%N)
	}
	//fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	//shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}})
	shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}})
}
