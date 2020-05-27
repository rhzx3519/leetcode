# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isSubtree(self, s, t):
        """
        :type s: TreeNode
        :type t: TreeNode
        :rtype: bool
        """
        def dfs(s, t):
            if s and t and s.val==t.val:
                return dfs(s.left, t.left) and dfs(s.right, t.right)
            if not s and not t: return True
            return False
        if not s and not t: return True
        if not s and t: return False
        return dfs(s, t) or self.isSubtree(s.left, t) or self.isSubtree(s.right, t)