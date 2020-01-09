# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        head = ListNode(0)
        p = head
        i,j = l1, l2
        while i and j:
            if i.val < j.val:
                p.next = i
                p = i
                i = i.next
            else:
                p.next = j
                p = j
                j = j.next
        while i:
            p.next = i
            p = i
            i = i.next
        while j:
            p.next = j
            p = j
            j = j.next
        
        t = head
        ret = head.next
        del t
        return ret
            
            
            
        