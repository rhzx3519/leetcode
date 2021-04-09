# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def delNodes(self, root, to_delete):
        """
        :type root: TreeNode
        :type to_delete: List[int]
        :rtype: List[TreeNode]
        """
        def post(root):
            if not root: return
            root.left = post(root.left)
            root.right = post(root.right)
            if root.val in to_delete:
                if root.left:
                    ans.append(root.left)
                if root.right:
                    ans.append(root.right)
                root = None
            return root
                    

        ans = []
        if root.val not in to_delete:
            ans.append(root)
        
        post(root)
        return ans

# 思路：后序遍历