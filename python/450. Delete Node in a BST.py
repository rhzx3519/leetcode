# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def deleteNode(self, root, key):
        """
        :type root: TreeNode
        :type key: int
        :rtype: TreeNode
        """
        if not root:
            return root

        if root.val > key:
            root.left = self.deleteNode(root.left, key)
        elif root.val < key:
            root.right = self.deleteNode(root.right, key)
        else:
            if not root.left or not root.right:
                root = root.left if root.left else root.right
            else:
                cur = root.right
                while cur.left:
                    cur = cur.left
                root.val = cur.val
                root.right = self.deleteNode(root.right, cur.val)

        return root


# 思路：

# 删除的节点有三种情况

# 删除的节点无子节点。无子节点直接删除该节点
# 删除的节点有且仅有一个子节点。子节点替换删除节点
# 删除的节点同时有两个节点。一种是从删除节点的右分支查找最小值，赋值给删除节点，再去删除找到的这个最小值；另一种是从删除节点的左分支查找最大值，赋值给删除节点，再去删除这个最大值。
# ————————————————
# 版权声明：本文为CSDN博主「diu_man」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
# 原文链接：https://blog.csdn.net/diu_man/article/details/81203522        