/*
*
@author ZhengHao Lou
Date    2022/12/26
*/
package main

import (
	"sort"
	"strings"
)

/*
*
https://leetcode.cn/problems/reward-top-k-students/
*/
func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	var pf, nf = make(map[string]bool), make(map[string]bool)
	for _, s := range positive_feedback {
		pf[s] = true
	}
	for _, s := range negative_feedback {
		nf[s] = true
	}
	var ids []int
	var scores = make(map[int]int)
	for i := range report {
		var c int
		for _, s := range strings.Split(report[i], " ") {
			if pf[s] {
				c += 3
			} else if nf[s] {
				c -= 1
			}
		}
		scores[student_id[i]] = c
		ids = append(ids, student_id[i])
	}

	sort.Slice(ids, func(i, j int) bool {
		if scores[ids[i]] != scores[ids[j]] {
			return scores[ids[i]] > scores[ids[j]]
		}
		return ids[i] < ids[j]
	})
	return ids[:k]
}
