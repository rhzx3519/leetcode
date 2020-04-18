# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        dummy = ListNode(0)
        p = dummy

        while head:
            tmp = head.next
            f = False
            while tmp and tmp.val==head.val:
                tmp = tmp.next
                f = True
            if not f:
                p.next = head
                p = p.next
                p.next = None                
            head = tmp
        return dummy.next