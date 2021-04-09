# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def subtreeWithAllDeepest(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if not root:
            return
        root.parent = None
        low = set()
        que = [root]
        while que:
            low = set()
            sz = len(que)
            while sz:
                sz -= 1
                root = que.pop(0)
                low.add(root)
                if root.left:
                    root.left.parent = root
                    que.append(root.left)
                if root.right:
                    root.right.parent = root
                    que.append(root.right)

        while len(low)!=1:        
            tmp = []
            for node in low:
                tmp.append(node.parent)
            low = set(tmp)
        return list(low)[0]
                