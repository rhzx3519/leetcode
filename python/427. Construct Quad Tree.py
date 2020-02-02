"""
# Definition for a QuadTree node.
class Node(object):
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight
"""
class Solution(object):
    def construct(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: Node
        """

        def dfs(x, y, offset):
            if offset<= 0: return

            for i in range(x, x+offset):
                for j in range(y, y+offset):
                    if grid[i][j] != grid[x][y]:
                        topLeft = dfs(x, y, offset/2)
                        topRight = dfs(x, y+offset/2, offset/2)
                        bottomLeft = dfs(x+offset/2, y, offset/2)
                        bottomRight = dfs(x+offset/2, y+offset/2, offset/2)
                        return Node(False, False, topLeft, topRight, bottomLeft, bottomRight)
            return Node(grid[x][y], True, None, None, None, None)

        return dfs(0, 0, len(grid))











