# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head or not head.next: return None
        f = False
        slow = head.next
        fast = head.next.next
        while fast and fast.next:
            if slow==fast:
                f = True
                break
            slow = slow.next
            fast = fast.next.next
        if not f: return None

        fast = head
        while fast != slow:
            slow = slow.next
            fast = fast.next
        return fast

if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(2)
    su = Solution()
    node = su.detectCycle(head)                    
    print node