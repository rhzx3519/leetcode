# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def maxLevelSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root: return 0
        lv = 0
        ans = 1
        max_val = float('-inf')
        que = [root]
        while que:
            s = 0
            lv += 1
            sz = len(que)
            while sz:
                root = que.pop(0)
                s += root.val
                if root.left:
                    que.append(root.left)
                if root.right:
                    que.append(root.right)
                sz -= 1
            
            # print lv, s
            if s > max_val:
                ans = lv
                max_val = s
        
        return ans