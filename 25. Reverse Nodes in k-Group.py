# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def reverse(self, head, tail):
        p = head
        new_head = ListNode(1)
        while p:
            t = p.next
            p.next = new_head.next
            new_head.next = p
            p = t
        t = new_head
        res = new_head.next
        del t
        return res
    
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        pre, p = None, head
        while p:
            q = p
            cnt = 1
            while q and cnt < k:
                q = q.next
                cnt += 1
            # print q.val, cnt 
            if q and cnt == k:
                t = q.next
                if pre == None:
                    q.next = None
                    head = self.reverse(p, q)
                else:
                    q.next= None
                    pre.next = self.reverse(p, q)
                pre = p
                p = t
            else:
                if pre:
                    pre.next = p
                break
        return head
                    
l1 = ListNode(1)
l1.next = ListNode(2)
l1.next.next = ListNode(3)
s = Solution()
p = s.reverseKGroup(l1, 1)
print p.val, p.next.val, p.next.next.val

        