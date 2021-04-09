# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def distributeCoins(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.ans = 0
        def dfs(root):
            if not root:
                return 0
            root.val += dfs(root.left)
            root.val += dfs(root.right)
            self.ans += abs(root.val - 1)
            return root.val - 1
        
        dfs(root)
        return self.ans