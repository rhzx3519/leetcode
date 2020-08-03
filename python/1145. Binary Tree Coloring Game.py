# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def btreeGameWinningMove(self, root, n, x):
        """
        :type root: TreeNode
        :type n: int
        :type x: int
        :rtype: bool
        """
        self.left = 0
        self.right = 0
        def count(root):
            if not root:
                return 0
            left = count(root.left)
            right = count(root.right)
            if (root.val==x):
                self.left, self.right = left, right
            return left + right + 1

        count(root)
        return n//2< max(self.left, self.right, n - self.left - self.right - 1)

# 思路：1st player选择一个节点之后，将树划分成三部分, x 的左右子树lef, right和x的父节点上面的树，n-left-right-1
# 只要保证总结点的一半<三部分中的最大值，就可以保证y获胜
