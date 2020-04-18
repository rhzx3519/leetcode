# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def pruneTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if not root: return root
        root.left = self.pruneTree(root.left)
        root.right = self.pruneTree(root.right)

        if not root.left and not root.right:
            return None if root.val==0 else root
    
        return root
