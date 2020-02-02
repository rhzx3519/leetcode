# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def findFrequentTreeSum(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        mp = collections.defaultdict(int)
        self.max_val = 0
        def dfs(root):
            if not root: return 0
            l = dfs(root.left)
            r = dfs(root.right)
            s = l + r + root.val
            mp[s] += 1
            self.max_val = max(self.max_val, mp[s])
            return s

        dfs(root)
        res = []
        return [k for k, v in mp.iteritems() if v==self.max_val]



