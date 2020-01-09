# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        pre, p = None, head
        while p and p.next:
            if pre == None:
                head = p.next
                t = p.next.next
                p.next = t
                head.next = p
            else:
                pre.next = p.next
                t = p.next.next
                p.next = t
                pre.next.next = p
            pre = p
            p = p.next
        return head