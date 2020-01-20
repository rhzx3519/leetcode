#!usr/bin/env python  
#-*- coding:utf-8 _*-  

from collections import deque

INF = 1<<30

class Solution(object):
    def networkDelayTime(self, times, N, K):
        """
        :type times: List[List[int]]
        :type N: int
        :type K: int
        :rtype: int
        """
        edges = [[INF]*(N+1) for _ in range(N+1)]
        vis = [0]*(N+1)
        for t in times:
            u = t[0]
            v = t[1]
            w = t[2]
            edges[u][v] = w

        # print edges

        dis = [INF]*(N+1)
        def dijkstra(K):
            # dis 存储K到其他节点的最短距离
            
            for i in range(1, N+1):
                dis[i] = edges[K][i]      
            
            vis[K] = 1
            for _ in range(1, N+1):

                min_idx = -1
                min_val = 1<<30
                for i in range(1, N+1):
                    if not vis[i] and dis[i] < min_val:
                        min_val = dis[i]
                        min_idx = i

                if min_idx==-1:
                    continue
                vis[min_idx] = 1
                for i in range(1, N+1):
                    if edges[min_idx][i] + dis[min_idx] < dis[i]:
                        dis[i] = edges[min_idx][i] + dis[min_idx]

        dijkstra(K)
        # print vis, dis
        if all(vis[1:]):
            max_val = 0
            for i in range(1, N+1):
                if i != K:
                    max_val = max(max_val, dis[i])
            return max_val

        return -1


if __name__ == '__main__':

    times = [[3,5,78],[2,1,1],[1,3,0],[4,3,59],[5,3,85],[5,2,22],[2,4,23],[1,4,43],[4,5,75],[5,1,15],[1,5,91],[4,1,16],[3,2,98],[3,4,22],[5,4,31],[1,2,0],[2,5,4],[4,2,51],[3,1,36],[2,3,59]]

    N = 5
    K = 5
    su = Solution()
    print su.networkDelayTime(times, N, K)


# 有 N 个网络节点，标记为 1 到 N。

# 给定一个列表 times，表示信号经过有向边的传递时间。 times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。

# 现在，我们向当前的节点 K 发送了一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。

# 注意:

# N 的范围在 [1, 100] 之间。
# K 的范围在 [1, N] 之间。
# times 的长度在 [1, 6000] 之间。
# 所有的边 times[i] = (u, v, w) 都有 1 <= u, v <= N 且 0 <= w <= 100。

# 来源：力扣（LeetCode）
# 链接：https://dev.lingkou.xyz/problems/network-delay-time
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。