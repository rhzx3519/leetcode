# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def generateTrees(self, n):
        """
        :type n: int
        :rtype: List[TreeNode]
        """
        if n==0: return []
        def dfs(start, end):
            res = []
            if start>end:
                res.append(None)
                return res
            
            for i in range(start, end+1):
                left_nodes = dfs(start, i-1)
                right_nodes = dfs(i+1, end)
                for l in left_nodes:
                    for r in right_nodes:
                        root = TreeNode(i)
                        root.left = l
                        root.right = r
                        res.append(root)
            return res

        return dfs(1, n)

# 思路：双索引的dfs        
