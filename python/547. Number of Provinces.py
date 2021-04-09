class Solution(object):
    def findCircleNum(self, isConnected):
        """
        :type isConnected: List[List[int]]
        :rtype: int
        """
        n = len(isConnected)
        vis = [0]*n

        def dfs(idx):
            vis[idx] = 1
            for i in range(n):
                if isConnected[idx][i] and not vis[i]:
                    dfs(i)
        
        ans = 0
        for i in range(n):
            if not vis[i]:
                dfs(i)
                ans += 1
        return ans