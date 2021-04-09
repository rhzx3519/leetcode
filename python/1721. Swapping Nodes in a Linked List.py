# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def swapNodes(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        t1 = t2 = None
        p = head
        cnt = 1
        while p:
            if cnt == k:
                t1 = p
                t2 = head
            elif cnt > k:
                t2 = t2.next
            p = p.next
            cnt += 1
        t1.val, t2.val = t2.val, t1.val
        return head
        
        # print t1.val, t2.val