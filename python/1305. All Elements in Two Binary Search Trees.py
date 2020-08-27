# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def getAllElements(self, root1, root2):
        """
        :type root1: TreeNode
        :type root2: TreeNode
        :rtype: List[int]
        """
        def inOrderTraverse(root, a):
            if not root:
                return
            inOrderTraverse(root.left, a)
            a.append(root.val)
            inOrderTraverse(root.right, a)

        a, b = [], []
        inOrderTraverse(root1, a)
        inOrderTraverse(root2, b)
        c = []
        i = j = 0
        na, nb = len(a), len(b)
        while i < na or j < nb:
            if i < na and j < nb:
                if a[i] < b[j]:
                    c.append(a[i])
                    i += 1
                else:
                    c.append(b[j])
                    j += 1
            elif i < na:
                c.append(a[i])
                i += 1
            else:
                c.append(b[j])
                j += 1
        return c