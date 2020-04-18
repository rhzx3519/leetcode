# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def sortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        def qsort(begin, end):  
            if begin!=end:
                p = partition(begin, end)
                qsort(begin, p)
                qsort(p.next, end)
                
        def partition(begin, end):
            p = begin
            q = begin
            key = p.val
            while p!=end:
                if p.val < key:
                    q = q.next
                    p.val, q.val = q.val, p.val
                p = p.next
            begin.val, q.val = q.val, begin.val
            return q

        qsort(head, None)
        return head


if __name__ == '__main__':
    l = ListNode(4)
    l.next = ListNode(2)
    l.next.next = ListNode(1)
    l.next.next = ListNode(3)

    su = Solution()
    su.sortList(l)        