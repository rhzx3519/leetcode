# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        res = []
        if not root:
            return res
        que = []
        que.append(root)
        while que:
            sz = len(que)
            a = []
            while sz:
                q = que.pop(0)
                a.append(q.val)
                if q.left:
                    que.append(q.left)
                if q.right:
                    que.append(q.right)
                sz -= 1
            res.append(a)
        
        return res