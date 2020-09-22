# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isSubPath(self, head, root):
        """
        :type head: ListNode
        :type root: TreeNode
        :rtype: bool
        """
        def dfs(head, root):
            if not head:
                return True
            if not root:
                return False
            if head.val == root.val:
                return dfs(head.next, root.left) or dfs(head.next, root.right)
            return False
        
        if not head:
            return True
        if not root:
            return False
        if dfs(head, root):
            return True
        return self.isSubPath(head, root.left) or self.isSubPath(head, root.right)



