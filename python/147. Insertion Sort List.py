# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def insertionSortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        dummy = ListNode(0)
        p = dummy
        while head:
            tmp = head
            head = head.next
            p = dummy
            while p.next and p.next.val <tmp.val:
                p = p.next
            k = p.next
            p.next = tmp
            tmp.next = k
        return dummy.next