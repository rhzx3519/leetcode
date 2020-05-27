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

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def sortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        return self.merge_sort(head)

    def merge_sort(self, head):
        if not head or not head.next:
            return head
        p = q = head
        pre = None
        while q and q.next:
            pre = p
            p = p.next
            q = q.next.next
        
        l = head
        r = pre.next
        pre.next = None
        l = self.merge_sort(l)
        r = self.merge_sort(r)
        return self.merge(l, r)
    
    def merge(self, p, q):
        dummy = ListNode(0)
        t = dummy
        while p or q:
            if p and q:
                if p.val < q.val:
                    t.next = p
                    p = p.next
                else:
                    t.next = q
                    q = q.next
            elif p:
                t.next = p
                p = p.next
            elif q:
                t.next = q
                q = q.next
            t = t.next
        return dummy.next
# 思路： 链表的归并排序，包括 1）分割点的确定，2）链表合并
# time: O(nlogn), space: O(1)





if __name__ == '__main__':
    l = ListNode(4)
    l.next = ListNode(2)
    l.next.next = ListNode(1)
    l.next.next = ListNode(3)

    su = Solution()
    su.sortList(l)        