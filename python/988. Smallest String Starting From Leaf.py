# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def smallestFromLeaf(self, root):
        """
        :type root: TreeNode
        :rtype: str
        """
        
        def dfs(root, a):
            if not root:
                return

            if not root.left and not root.right:
                s = ''.join([chr(i + ord('a')) for i in a])
                self.res = min(self.res, s[::-1])
                return

            if root.left:
                dfs(root.left, a + [root.left.val])
            if root.right: 
                dfs(root.right, a + [root.right.val])

        self.res = 'z'*10001
        dfs(root, [root.val])
        return self.res

        