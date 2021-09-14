package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1, s2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i := 0; i < max(len(s1), len(s2)); i++ {
		var t1, t2 int
		if i < len(s1) {
			t1, _ = strconv.Atoi(s1[i])
		}
		if i < len(s2) {
			t2, _ = strconv.Atoi(s2[i])
		}
		if t1 < t2 {
			return -1
		} else if t1 > t2 {
			return 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
