package main

import "sort"

/*
*
https://leetcode.cn/problems/high-access-employees/
*/
func findHighAccessEmployees(access_times [][]string) []string {
	var access = make(map[string][]int)
	for _, t := range access_times {
		access[t[0]] = append(access[t[0]], convert(t[1]))
	}

	var ans []string
	for x, a := range access {
		sort.Ints(a)
		for i := 2; i < len(a); i++ {
			if a[i]-a[i-2] < 60 {
				ans = append(ans, x)
				break
			}
		}
	}
	return ans
}

func convert(t string) int {
	return (10*int(t[0]-'0')+int(t[1]-'0'))*60 + int(t[2]-'0')*10 + int(t[3]-'0')
}
