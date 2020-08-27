class Solution(object):
    def makeConnected(self, n, connections):
        """
        :type n: int
        :type connections: List[List[int]]
        :rtype: int
        """
        nc = len(connections)
        if nc < n-1:
            return -1
        parent = [i for i in range(n)]

        def find(x):
            px = parent[x]
            if px==x:
                return px
            return find(px)

        def union(x, y):
            px = find(x)
            py = find(y)
            parent[py] = px
        
        for a, b in connections:
            union(a, b)
        
        count = 0
        for i, num in enumerate(parent):
            if i==num:
                count += 1
        return count -1

# 思路：n个节点最少可以通过n-1条线连接
# 计算图中的连通块个数C, 让所有节点联通的线个数 为C-1
# time: O(E), space: O(N)        