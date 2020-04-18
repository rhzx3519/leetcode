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
            p.next = head
            p = p.next
            tmp = head.next
            while tmp and tmp.val==head.val:
                tmp = tmp.next
            head = tmp

        p.next = None
        return dummy.next