class Solution(object):
    def minTime(self, n, edges, hasApple):
        """
        :type n: int
        :type edges: List[List[int]]
        :type hasApple: List[bool]
        :rtype: int
        """
        parent = [-1] * n
        def getParent(i):
            if parent[i]==-1:
                return -1
            elif parent[i]==0:
                return 0
            return getParent(parent[i])
        
        for a, b in edges:
            if getParent(b)==0:
                parent[a] = b
            else:
                parent[b] = a
        
        # print parent
        ans = 0
        vis = [0] * n
        vis[0] = 1
        for i in range(1, n):
            if not hasApple[i]:
                continue
            t = i
            while not vis[t]:
                vis[t] = 1
                ans += 1
                t = parent[t]
        return ans*2