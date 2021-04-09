# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class CBTInserter(object):

    def __init__(self, root):
        """
        :type root: TreeNode
        """
        self.root = root
        self.nodes = []
        que = [root]
        while que:
            root = que.pop(0)
            self.nodes.append(root)
            if root.left:
                que.append(root.left)
            if root.right:
                que.append(root.right)
        

    def insert(self, v):
        """
        :type v: int
        :rtype: int
        """
        self.nodes.append(TreeNode(v))
        n = len(self.nodes)
        idx = 0
        if n&1:
            idx = (n-1)/2 - 1
            self.nodes[idx].right = self.nodes[-1]
        else:
            idx = n/2 - 1
            self.nodes[idx].left = self.nodes[-1]
        return self.nodes[idx].val

    def get_root(self):
        """
        :rtype: TreeNode
        """
        return self.root
        


# Your CBTInserter object will be instantiated and called as such:
# obj = CBTInserter(root)
# param_1 = obj.insert(v)
# param_2 = obj.get_root()