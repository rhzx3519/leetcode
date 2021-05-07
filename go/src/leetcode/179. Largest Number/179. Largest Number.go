package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, 0)
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] + strs[j] > strs[j] + strs[i]
	})

	s := strings.Join(strs, "")
	i := 0
	for ; i < len(s) && s[i] == '0'; i++ {}
	if i == len(s) {
		return "0"
	}
	return s[i:]
}