# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def lcaDeepestLeaves(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        def depth(root):
            if not root: return 0
            l = depth(root.left)
            r = depth(root.right)
            return max(l, r) + 1

        if not root: return
        l = depth(root.left)
        r = depth(root.right)
        if l==r:
            return root
        elif l > r:
            return self.lcaDeepestLeaves(root.left)
        else:
            return self.lcaDeepestLeaves(root.right)
