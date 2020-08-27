# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def deepestLeavesSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        res = 0
        que = [root]
        while que:
            sz = len(que)
            res = 0
            while sz:
                sz -= 1
                root = que.pop(0)
                if root.left:
                    que.append(root.left)
                if root.right:
                    que.append(root.right)
                res += root.val
        return res
            
