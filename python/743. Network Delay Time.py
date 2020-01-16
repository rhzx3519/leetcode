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