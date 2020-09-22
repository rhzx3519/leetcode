# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def longestZigZag(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        self.ans = 0
        def dfs(node, isLeft):
            if not node:
                return 0
            l = dfs(node.left, True)
            r = dfs(node.right, False)
            self.ans = max(self.ans, l, r)
            return r + 1 if isLeft else l + 1

        dfs(root, False)
        return self.ans    
