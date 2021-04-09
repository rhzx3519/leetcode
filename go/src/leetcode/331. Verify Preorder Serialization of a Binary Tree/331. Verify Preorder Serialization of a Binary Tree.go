package main

import "strings"

func isValidSerialization(preorder string) bool {
	diff := 1
	nodes := strings.Split(preorder, ",")
	for _, node := range nodes {
		diff--
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff==0
}
