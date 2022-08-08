/**
@author ZhengHao Lou
Date    2022/05/31
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/
方法二：0-1 广度优先搜索
常规的广度优先搜索可以找出在边权均为 11 时的单源最短路，然而在我们的建模中，边权除了 11 之外也可能为 00。我们是否可以修改广度优先搜索的算法框架，使得它可以找出在边权为 00 或 11 时的单源最短路呢？

答案是可以的。这种修改过的广度优先搜索被称为「0-1 广度优先搜索」，这里 有一篇很详细的教程。

保证广度优先搜索正确性的基础在于：对于源点 ss 以及任意两个节点 uu 和 vv，如果 \textit{dist}(s, u) < \textit{dist}(s, v)dist(s,u)<dist(s,v)（其中 \textit{dist}(x, y)dist(x,y) 表示从节点 xx 到节点 yy 的最短路长度），那么节点 uu 一定会比节点 vv 先被取出队列。在常规的广度优先搜索中，我们使用队列作为维护节点的数据结构，就保证了从队列中取出的节点，它们与源点之间的距离是单调递增的。然而如果边权可能为 00，就会出现如下的情况：

源点 ss 被取出队列；

源点 ss 到节点 v_1v
1
​
  有一条权值为 11 的边，将节点 v_1v
1
​
  加入队列；

源点 ss 到节点 v_2v
2
​
  有一条权值为 00 的边，将节点 v_2v
2
​
  加入队列；

此时节点 v_2v
2
​
  一定会在节点 v_1v
1
​
  之后被取出队列，但节点 v_2v
2
​
  与源点之间的距离反而较小，这样就破坏了广度优先搜索正确性的基础。

那么我们如何修改广度优先搜索的算法框架呢？我们可以使用双端队列（double-ended queue, deque）代替普通的队列作为维护节点的数据结构。当任一节点 uu 被取出队列时，如果它到某节点 v_iv
i
​
  有一条权值为 00 的边，那么就将节点 v_iv
i
​
  加入双端队列的「队首」。如果它到某节点 v_jv
j
​
  有一条权值为 11 的边，那么和常规的广度优先搜索相同，我们将节点 v_jv
j
​
  加入双端队列的「队尾」。这样以来，我们保证了任意时刻从队首到队尾的所有节点，它们与源点之间的距离是单调递增的，即从队列中取出的节点与源点之间的距离同样是单调递增的。

0-1 广度优先搜索的实现其实与 Dijkstra 算法非常相似。在 Dijkstra 算法中，我们用优先队列保证了距离的单调递增性。而在 0-1 广度优先搜索中，实际上任意时刻队列中的节点与源点的距离均为 dd 或 d + 1d+1（其中 dd 为某一非负整数），并且所有与源点距离为 dd 的节点都出现在队首附近，所有与源点距离为 d + 1d+1 的节点都出现在队尾附近。因此，我们只要使用双端队列，对于边权为 00 和 11 的两种情况分别将对应节点添加至队首和队尾，就保证了距离的单调递增性。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/solution/shi-wang-ge-tu-zhi-shao-you-yi-tiao-you-xiao-lu-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type pair struct {
	x, y int
}

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	INF := m * n
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = INF
		}
	}

	dis[0][0] = 0
	var que = []pair{{}}
	for len(que) != 0 {
		p := que[0]
		que = que[1:]
		g := grid[p.x][p.y]
		for _, dir := range dirs {
			nx := p.x + dir[0]
			ny := p.y + dir[1]
			if nx < 0 || nx == m || ny < 0 || ny == n {
				continue
			}
			if dis[p.x][p.y]+g < dis[nx][ny] {
				dis[nx][ny] = dis[p.x][p.y] + g
				if grid[nx][ny] == 0 {
					que = append(que, pair{})
					copy(que[1:], que[0:])
					que[0] = pair{x: nx, y: ny}
				} else {
					que = append(que, pair{x: nx, y: ny})
				}
			}
		}
	}

	fmt.Println(dis)
	return dis[m-1][n-1]
}

func main() {
	minimumObstacles([][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}})
	minimumObstacles([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}})
}
