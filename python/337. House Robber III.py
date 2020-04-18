# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def rob(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        def dfs(root):
            res = [0]*2  # res[0]不取当前结点，res[1]取当前结点
            if not root:
                return res
            lres = dfs(root.left)
            rres = dfs(root.right)
            res[0] = max(lres[0], lres[1]) + max(rres[0], rres[1])
            res[1] = root.val + lres[0] + rres[0]
            return res

        res = dfs(root)
        return max(res)




