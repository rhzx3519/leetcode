# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isEvenOddTree(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        que = [root]
        f = 1 
        while que:
            sz = len(que)
            pre = float('-inf') if f else float('inf')
            while sz:
                sz -= 1
                root = que.pop(0)
                val = root.val
                if f:
                    if val&1 == 0:
                        return False
                    if val <= pre:
                        return False
                else:
                    if val&1 == 1:
                        return False
                    if val >= pre:
                        return False
                pre = val
                if root.left:
                    que.append(root.left)
                if root.right:
                    que.append(root.right)
            f ^= 1

        return True
        