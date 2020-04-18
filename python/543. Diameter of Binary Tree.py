# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def diameterOfBinaryTree(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.res = 0
        def dfs(node):
            if not node: return 0
            l = dfs(node.left)
            r = dfs(node.right)
            self.res = max(self.res, l+r)
            return max(l, r) + 1

        dfs(root)
        return self.res

# 思路，求左子树深度+右子树深度和的最大值        
            