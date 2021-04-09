class Solution(object):
    def maxProbability(self, n, edges, succProb, start, end):
        """
        :type n: int
        :type edges: List[List[int]]
        :type succProb: List[float]
        :type start: int
        :type end: int
        :rtype: float
        """
        es = [[] for _ in range(n)]
        for i, (a, b) in enumerate(edges):
            es[a].append((b, succProb[i]))
            es[b].append((a, succProb[i]))
        
        vis = [0]*n
        dis = [0]*n
        def dijkstra(K, n):
            # dis 存储K到其他节点的最短距离
            
            for nx, p in es[K]:
                dis[nx] = p
            
            vis[K] = 1
            for _ in range(n):

                max_idx = -1
                max_val = 0
                for i in range(n):
                    if not vis[i] and dis[i] > max_val:
                        max_val = dis[i]
                        max_idx = i

                if max_idx==-1:
                    continue
                vis[max_idx] = 1
                for nx, p in es[max_idx]:
                    if p*max_val > dis[nx]:
                        dis[nx] = p*max_val
                # for i in range(n):
                #     if es[max_idx][i] * dis[max_idx] > dis[i]:
                #         dis[i] = es[max_idx][i] * dis[max_idx]
        
        dijkstra(start, n)
        return dis[end]