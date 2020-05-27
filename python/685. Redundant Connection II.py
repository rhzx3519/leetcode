class Solution(object):
    def findRedundantDirectedConnection(self, edges):
        """
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        n = len(edges)
        look_up = [i for i in range(n+1)]
        
        def find(x):
            if look_up[x]!=x:
                look_up[x] = find(look_up[x])
            return look_up[x]

        def union(x, y):
            px = find(x)
            py = find(y)
            if px!=py:
                look_up[py] = px
        
        for x, y in edges:
            px = find(x)
            py = find(y)
            if py!=y and py!=px:
                return (py, y)
            else:
                union(x, y)
        return edges[-1]

if __name__ == '__main__':
    edges = [[1,2], [2,3], [3,4], [4,1], [1,5]]
    # edges = [[1,2], [1,3], [2,3]]
    # edges = [[2,1],[3,1],[4,2],[1,4]]
    su = Solution()
    print su.findRedundantDirectedConnection(edges)        