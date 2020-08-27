# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def goodNodes(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.count = 0

        def traverse(root, curMax):
            if not root:
                return
            if curMax <= root.val:
                self.count += 1
            traverse(root.left, max(curMax, root.val))
            traverse(root.right, max(curMax, root.val))
        
        traverse(root, -float('inf'))
        return self.count