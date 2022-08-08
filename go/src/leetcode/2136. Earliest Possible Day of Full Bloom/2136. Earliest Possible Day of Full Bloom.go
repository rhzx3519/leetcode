/**
@author ZhengHao Lou
Date    2022/01/11
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/earliest-possible-day-of-full-bloom/
其实这题就是利用的《操作系统》中的《优先级调度算法》的思想（发散一下408思维）。

1）把 播种时间 看作 进程使用cpu的时间 ，把 生长时间 看作 I/O设备工作的时间，由于I/O型设备的处理速度比CPU慢得多，只有让I/O设备尽早地进行工作，cpu才能调度其他进程执行，这样会提升系统的整体效率。
2）这样转换来看，肯定要cpu先调度I/O时间长的进程先执行，才能在I/O设备工作的时候尽可能的调度其他程序执行，这样最终的时间会更短。对应本题的话，就是让生长时间长的花先播种，这样在该花生长的时候，可以尽可能地播种其他花。
3）对于最晚时间的理解，可以画下甘特图，这样会更加清晰。
 */
func earliestFullBloom(plantTime []int, growTime []int) int {
	tasks := []int{}
	n :=  len(plantTime)
	for i := 0; i < n; i++ {
		tasks = append(tasks, i)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return growTime[tasks[i]] > growTime[tasks[j]]
	})
	var total, day int
	for _, i := range tasks {
		day += plantTime[i]
		total = max(total, day + growTime[i])
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(earliestFullBloom([]int{1,4,3}, []int{2,3,1}))
}
