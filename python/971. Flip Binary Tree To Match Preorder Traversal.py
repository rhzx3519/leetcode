# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def flipMatchVoyage(self, root, voyage):
        """
        :type root: TreeNode
        :type voyage: List[int]
        :rtype: List[int]
        """
        n = len(voyage)
        self.idx = 0
        self.res = []
        def dfs(root):
            if not root: return
            if root.val != voyage[self.idx]:
                self.res = [-1]
                return

            self.idx += 1
            if self.idx<n and root.left and root.left.val!=voyage[self.idx]:
                self.res.append(root.val)
                dfs(root.right)
                dfs(root.left)
            else:
                dfs(root.left)
                dfs(root.right)

        dfs(root)
        return self.res

if __name__ == '__main__':
    node = TreeNode(1)                    
    node.left = TreeNode(2)
    node.right = TreeNode(3)
    voyage = [1, 2, 3]
    su = Solution()
    print su.flipMatchVoyage(node, voyage)