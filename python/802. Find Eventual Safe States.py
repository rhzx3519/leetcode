class Solution(object):
    def eventualSafeNodes(self, graph):
        """
        :type graph: List[List[int]]
        :rtype: List[int]
        """
        n = len(graph)
        vis = [0]*n 
        def dfs(cur):
            if vis[cur]==-1:
                return False
            if vis[cur]==1:
                return True
            vis[cur] = -1
            for i in graph[cur]:
                if not dfs(i):
                    return False
            vis[cur] = 1
            return True
        
        for i in range(n):
            if not vis[i]:
                dfs(i)
        return [i for i in range(n) if vis[i]==1]

# 寻找不在环中的结点