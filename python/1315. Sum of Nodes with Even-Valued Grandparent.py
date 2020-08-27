# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def sumEvenGrandparent(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        ans = 0
        f = root.val&1
        if not f:  
            if root.left:
                if root.left.left:
                    ans += root.left.left.val
                if root.left.right:
                    ans += root.left.right.val
            if root.right:
                if root.right.left:
                    ans += root.right.left.val
                if root.right.right:
                    ans += root.right.right.val

        ans += self.sumEvenGrandparent(root.left) + self.sumEvenGrandparent(root.right)
        return ans


        