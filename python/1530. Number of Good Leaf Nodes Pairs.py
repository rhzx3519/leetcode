# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def countPairs(self, root, distance):
        """
        :type root: TreeNode
        :type distance: int
        :rtype: int
        """
        self.ans = 0

        def post(root):
            if not root: return []
            if not root.left and not root.right:
                return [0]
            ret = []
            left = post(root.left)
            for i, x in enumerate(left):
                if x+1 <= distance:
                    ret.append(x+1)
                    left[i] += 1

            right = post(root.right)
            for i, x in enumerate(right):
                if x+1 <= distance:
                    ret.append(x+1)
                    right[i] += 1      

            for l in left:
                for r in right:
                    if l+r <= distance:
                        self.ans += 1      

            return ret

        post(root)
        return self.ans
        