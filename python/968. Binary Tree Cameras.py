# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def minCameraCover(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.ans = 0
        if not root:
            return self.ans
        
        # 0：该节点安装了监视器 1：该节点可观，但没有安装监视器 2：该节点不可观
        def dfs(root):
            if not root:
                return 1
            l, r = dfs(root.left), dfs(root.right)
            if l==2 or r==2:
                self.ans += 1 # 左右孩子有一个不可观测，需要在root安装监控器
                return 0
            elif l==0 or r==0: # 左右孩子有一个安装了监控器，root变成状态1：可观测且没有安装监控器
                return 1
            return 2 # root不可观测

        if dfs(root)==2:
            self.ans += 1
        return self.ans