# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def bstToGst(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        def dfs(root, val):
            if not root: return val
            right = dfs(root.right, val)
            root.val += right
            left = dfs(root.left, root.val)
            return left

        dfs(root, 0)
        return root
