class Solution(object):
    def maximalNetworkRank(self, n, roads):
        """
        :type n: int
        :type roads: List[List[int]]
        :rtype: int
        """
        ans = 0
        edges = [[] for _ in range(n)]
        for a, b in roads:
            edges[a].append(b)
            edges[b].append(a)
            
        for i in range(n):
            for j in range(i):
                t = len(edges[i]) + len(edges[j])
                t -= int(i in edges[j])
                ans = max(ans, t)
        return ans


