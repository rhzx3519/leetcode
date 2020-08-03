class Solution(object):
    def loudAndRich(self, richer, quiet):
        """
        :type richer: List[List[int]]
        :type quiet: List[int]
        :rtype: List[int]
        """
        N = len(quiet)
        graph = [[] for _ in range(N)]
        for u, v in richer:
            graph[v].append(u)
        
        mem = {}
        def dfs(node):
            if node in mem:
                return mem[node]
            ans = node
            for child in graph[node]:
                cand = dfs(child)
                if quiet[cand] < quiet[ans]:
                    ans = cand
            mem[node] = ans
            return ans
        
        return map(dfs, range(N))

# 思路：创建一个有向图，graph[x] ->  y, y 的财富大于x，dfs遍历，使用备忘录剪枝
# time: O(N^2), space: O(N)