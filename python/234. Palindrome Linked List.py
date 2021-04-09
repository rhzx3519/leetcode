# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def isPalindrome(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        if not head:
            return True
        cnt = 0
        dummy = ListNode(0)
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            cnt += 1
        
        p = slow
        tail = p
        while p:
            tmp = p.next
            p.next = dummy.next
            dummy.next = p
            p = tmp
            cnt += 1
        tail.next = None
        # print cnt
        # if cnt&1:
        p = dummy.next
        while p and head:
            if p.val != head.val:
                return False
            p = p.next
            head = head.next
        return True
