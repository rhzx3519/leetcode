# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def partition(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """
        p1 = dummy1 = ListNode(0)
        p2 = dummy2 = ListNode(0)
        while head:
            if head.val < x:
                p1.next = head
                p1 = p1.next
            else:
                p2.next = head
                p2 = p2.next
            head = head.next
            
        p2.next = None
        p1.next = dummy2.next
        return dummy1.next

def buildList(a):
    if not a: return
    dummy = ListNode(0)
    p = dummy
    for i in a:
        p.next = ListNode(i)
        p = p.next
    return dummy.next        

def printList(head):
    a = []
    while head:
        a.append(head.val)
        head =head.next
    return '->'.join(map(str, a))    

if __name__ == '__main__':
    head = buildList([1,4,3,2,5,2])
    su = Solution()
    res = su.partition(head, 3)                     
    print(printList(res))