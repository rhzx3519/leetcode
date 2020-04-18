# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def widthOfBinaryTree(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.max_width = 0
        self.dfs(root, 1, 1, [])
        return self.max_width

    def dfs(self, root, idx, level, a):
        if not root:
            return
        if level > len(a):
            a.append(idx)
        self.max_width = max(self.max_width, idx - a[level-1] + 1)
        self.dfs(root.left, 2*idx, level + 1, a)
        self.dfs(root.right, 2*idx+1, level + 1, a)


## 数组a存储每层最左边的节点下标，dfs遍历二叉树时，当当前深度level比数组a大时，当前节点便是当前层最左侧的节点        