package tool

import (
	"math"
	"sort"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func BuildList(a []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, num := range a {
		p.Next = &ListNode{num, nil}
		p = p.Next
	}
	return dummy.Next
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SortString(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j+1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


