# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def __init__(self):
        self.last = -1<<64

    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root: return True

        if self.isValidBST(root.left):
            if self.last < root.val:
                self.last = root.val
                return self.isValidBST(root.right)
        return False