# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def maxAncestorDiff(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.res = -(1<<31)
        max_val = -(1<<31)
        min_val = 1<<31-1

        def dfs(root, max_val, min_val):
            if not root: return
            max_val = max(max_val, root.val)
            min_val = min(min_val, root.val)
            self.res = max(self.res, max(abs(max_val - root.val), abs(min_val - root.val)))
            dfs(root.left, max_val, min_val)
            dfs(root.right, max_val, min_val)
            
        dfs(root, max_val, min_val)
        return self.res