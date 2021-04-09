# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isCompleteTree(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return False
        que = [root]

        while que:
            root = que.pop(0)
            if not root:
                return not filter(lambda x: x, que)
            que.append(root.left)
            que.append(root.right)

        return True