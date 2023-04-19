package main

import (
	"fmt"
)

/**
1.(Medium)
某幼儿园组织小朋友们去儿童游乐园玩耍。如果你问某个小朋友班里有多少人，那么他会告诉你他所在的班级里除了他自己外有多少人，
将得到回答的人数放入members[]数组中, 返回游乐园中小朋友可能的最少人数

说明:
- members 数组的长度最大为1000
- members[i] 是在 [0, 999] 范围内的整数

示例:
输入: members = [1, 1, 2, 1]
输出: 7
解释:
(1)两个回答了 "1" 的小朋友可能有来自同一班级，设为一班。
(2)之后回答了 "2" 的小朋友不会是一班的，否则他们的回答会相互矛盾。设回答了 "2" 的小朋友为二班。
(3) 若(1)成立, 则最后一个回答1的小朋友不可能来自一班，否则1班将会有3人。设该小朋友来自三班
此外，还应有另外 2 个二班和 1 个三班小朋友的回答没有包含在数组中。
因此游乐园中小朋友最少7位 ： 4 位回答的和 3 位没有回答的

 */
func minCount(members []int) int {
	mapper := make(map[int][]int)
	for _, m := range members {
		if clz, ok := mapper[m]; ok {
			num := clz[len(clz) - 1]
			if num == m+1 {
				mapper[m] = append(mapper[m], 1)
			} else {
				mapper[m][len(mapper[m]) - 1]++
			}
		} else {
			mapper[m] = append(mapper[m], 1)
		}
	}
	var count int
	for m, clazz := range mapper {
		for _ = range clazz {
			count += m+1
		}
	}

	fmt.Println(count)
	return count
}

func main() {
	minCount([]int{1,1,2,1})		// 7
	minCount([]int{10,10,10,10})	// 11
	minCount([]int{})				// 0
	minCount([]int{0,0})			// 2
	minCount([]int{0,0,1,1,1})		// 6
	minCount([]int{0,2,2,2,4,2})	// 12
}
