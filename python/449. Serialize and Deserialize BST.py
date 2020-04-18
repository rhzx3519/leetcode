# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        def pre_order(root):
            out = []
            if root:
                out.append(str(root.val))
                out += pre_order(root.left)
                out += pre_order(root.right)
            return out
        return ','.join(pre_order(root))

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if not data:
            return None

        def build_tree(a, b):
            if not a:
                return None
            mid = a[0]
            root = TreeNode(mid)
            i = b.index(mid)
            root.left = build_tree(a[1: i+1], b[:i])
            root.right = build_tree(a[i+1:], b[i+1:])
            return root
        a = list(map(int, data.split(',')))
        b = sorted(a)
        return build_tree(a, b)


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))