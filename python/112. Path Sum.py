# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def hasPathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: bool
        """
        if not root: return False
        if not root.left and not root.right:
            return sum==root.val
        if root.left and root.right:
            return self.hasPathSum(root.left, sum-root.val) or self.hasPathSum(root.right, sum-root.val)
        if root.left:
            return self.hasPathSum(root.left, sum-root.val)
        if root.right:
            return self.hasPathSum(root.right, sum-root.val)
# 思路： dfs，叶子节点判断节点值是否为sum，非叶子节点判断递归判断子节点
# 时间复杂度:O(N), 空间复杂度O(N)           