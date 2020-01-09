# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        i = 0
        p = q = head
        while p:
            i += 1
            p = p.next
        i = i - n
        p, pre = head, None
        while i:
            pre = p
            p = p.next
            i -= 1
        if pre and p:
            t = p
            pre.next = p.next
            del t
            return head
            
        t = head
        head = head.next
        del t
        
        return head