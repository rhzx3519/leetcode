# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def findMode(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        ans = []
        max_val = float('-inf')
        cnt = collections.defaultdict(int)
        st = []
        while root or st:
            if root:
                st.append(root)
                root = root.left
            else:
                root = st.pop()                    
                cnt[root.val] += 1
                if cnt[root.val] > max_val:
                    max_val = cnt[root.val]
                    ans = [root.val]
                elif cnt[root.val] == max_val:
                    ans.append(root.val)
                root = root.right
        return ans