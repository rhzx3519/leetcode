# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def distanceK(self, root, target, K):
        """
        :type root: TreeNode
        :type target: TreeNode
        :type K: int
        :rtype: List[int]
        """
        def dfs(root, parent):
            if not root:
                return
            root.parent = parent
            dfs(root.left, root)
            dfs(root.right, root)
        
        dfs(root, None)

        K += 1
        visit = set()
        ans = []
        que = [target]
        while que:
            tmp = []
            K -= 1
            sz = len(que)
            while sz:
                sz -= 1
                root = que.pop(0)
                visit.add(root)
                tmp.append(root.val)
                if root.left and root.left not in visit:
                    que.append(root.left)
                if root.right and root.right not in visit:
                    que.append(root.right)
                if root.parent and root.parent not in visit:
                    que.append(root.parent)
            if K==0:
                ans = tmp
                break
        return ans
                

            