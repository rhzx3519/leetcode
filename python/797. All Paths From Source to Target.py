class Solution(object):
    def allPathsSourceTarget(self, graph):
        """
        :type graph: List[List[int]]
        :rtype: List[List[int]]
        """
        n = len(graph)
        vis = [0]*n
        res = []
        def dfs(cur, path):
            if cur==n-1:
                res.append(path)
                return

            for nx in graph[cur]:
                if vis[nx]: continue
                vis[nx] = 1
                dfs(nx, path+[nx])
                vis[nx] = 0

        dfs(0, [0])
        return res