# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def insertIntoMaxTree(self, root, val):
        """
        :type root: TreeNode
        :type val: int
        :rtype: TreeNode
        """
        node = TreeNode(val)
        if not root or val > root.val:
            node.left = root
            return node
        
        tmp = root
        while tmp.right and tmp.right.val > val:
            tmp = tmp.right
        node.left = tmp.right
        tmp.right = node
        return root

# 附加值添加到A的末尾        