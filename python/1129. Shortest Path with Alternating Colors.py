class Solution(object):
    def shortestAlternatingPaths(self, n, red_edges, blue_edges):
        """
        :type n: int
        :type red_edges: List[List[int]]
        :type blue_edges: List[List[int]]
        :rtype: List[int]
        """
        redEdges = [[] for _ in range(n)]
        blueEdges = [[] for _ in range(n)]
        for a, b in red_edges:
            redEdges[a].append(b)
        for a, b in blue_edges:
            blueEdges[a].append(b)

        
        def bfs(color):
            # red: 0, blue: 1
            ans = [float('inf')]*n
            ans[0] = 0
            redVis = [0]*n
            blueVis = [0]*n
            lv = 0
            que = [0]
            while que:
                lv += 1
                sz = len(que)
                while sz:
                    sz -= 1
                    node = que.pop(0)
                    edges = redEdges if color==0 else blueEdges
                    vis = redVis if color==0 else blueVis
                    for i, nx in enumerate(edges[node]):
                        if vis[nx]==1: continue
                        vis[nx] = 1
                        if ans[nx]==float('inf'): ans[nx] = lv
                        que.append(nx)
                color ^= 1
            return ans
        
        ans1 = bfs(0)
        ans2 = bfs(1)
        ans = [-1]*n
        for i in range(n):
            ans[i] = -1 if min(ans1[i], ans2[i])==float('inf') else min(ans1[i], ans2[i])
        return ans

# 思路：使用两个vis数组分别记录红色边访问和蓝色边访问


