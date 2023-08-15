package main

import "fmt"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	//k := len(indices)
	n := len(s)
	f := make([]int, n)
	for i := range f {
		f[i] = -1
	}
	for i, j := range indices {
		l := len(sources[i])
		if j+l <= n && s[j:j+l] == sources[i] {
			f[j] = i
		}
	}
	var bytes []byte
	for i := 0; i < n; {
		if f[i] != -1 {
			bytes = append(bytes, targets[f[i]]...)
			i += len(sources[f[i]])
		} else {
			bytes = append(bytes, s[i])
			i++
		}
	}
	return string(bytes)
}

func main() {
	fmt.Printf(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}
