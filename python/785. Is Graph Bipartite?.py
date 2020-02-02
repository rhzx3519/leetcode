class Solution(object):
    def isBipartite(self, graph):
        """
        :type graph: List[List[int]]
        :rtype: bool
        """
        n = len(graph)
        colors = [0]*n

        def dfs(idx, color):
            if colors[idx]!=0:
                return colors[idx]==color
            colors[idx] = color
            for i in graph[idx]:
                if not dfs(i, -color):
                    return False
            return True

        for i in range(n):
            if colors[i]==0 and not dfs(i, 1):
                return False
        return True


if __name__ == '__main__':
    graph = [[1,3],[0,2],[1,3],[0,2]]
    su = Solution()
    su.isBipartite(graph)