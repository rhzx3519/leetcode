class Solution(object):
    def minCostConnectPoints(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        def dis(p1, p2):
            return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

        if not points:
            return 0
        n = len(points)
        grid = [[0]*n for _ in range(n)]
        for i in range(n):
            for j in range(i+1, n):
                d = dis(points[i], points[j])
                grid[i][j] = d
                grid[j][i] = d
        # print grid
        ans = 0
        dis = [float('inf')]*n # 最小生成树到剩余节点的距离
        vis = [0]*n
        for i in range(n):
            t = -1
            for j in range(n):
                if vis[j]:
                    continue
                if t==-1 or dis[j] < dis[t]:
                    t = j
            vis[t] = 1
            if i: ans += dis[t]
            for j in range(n):
                if not vis[j]:
                    dis[j] = min(dis[j], grid[t][j])
        return ans
                
# 思路：最小生成树，Prim算法，          
# time: O(N^2), space: O(N^2)  
            
