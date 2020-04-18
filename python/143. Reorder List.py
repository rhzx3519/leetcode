# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def reorderList(self, head):
        """
        :type head: ListNode
        :rtype: None Do not return anything, modify head in-place instead.
        """
        if not head or not head.next: return
        p = head
        st = []
        while p:
            st.append(p)
            p = p.next
        
        dummy = ListNode(0)
        p = dummy
        while len(st)>1:
            p.next = st.pop(0)
            p = p.next
            print p.val
            p.next = st.pop()
            p = p.next
            print p.val

        if st:
            p.next = st[-1]
            p = p.next
        
        p.next = None

            
if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)
    head.next.next.next = ListNode(4)
    head.next.next.next.next = ListNode(5)
    su = Solution()
    su.reorderList(head)                            