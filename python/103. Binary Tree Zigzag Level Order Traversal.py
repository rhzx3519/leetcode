# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def zigzagLevelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        ans = []
        if not root: return ans
        que = [root]
        f = 1
        while que:
            sz = len(que)
            a = []
            while sz:
                sz -= 1
                root = que.pop(0)
                a.append(root.val)
                if root.left:
                    que.append(root.left)
                if root.right:
                    que.append(root.right)
            if f:
                ans.append(a)
            else:
                a.reverse()
                ans.append(a)
            f = 1-f
        return ans