#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return '#'
        return '{},{},{}'.format(root.val, self.serialize(root.left), self.serialize(root.right))

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        a = data.split(',')

        def dfs(a):
            if a[0]=='#':
                a.pop(0)
                return
            root = TreeNode(int(a[0]))
            a.pop(0)
            root.left = dfs(a)
            root.right = dfs(a)
            return root

        return dfs(a)


if __name__ == '__main__':
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.right.left = TreeNode(4)
    root.right.right = TreeNode(5)

    codec = Codec()
    print codec.serialize(root)
    root1 = codec.deserialize(codec.serialize(root))
    print codec.serialize(root1)

    


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))