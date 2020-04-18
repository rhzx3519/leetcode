# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseBetween(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """
        st = []
        dummy = ListNode(0)
        p = dummy
        p.next = head
        idx = 1
        while head and idx!=m:
            p = head
            head = head.next
            idx += 1

        tmp = head
        while tmp and idx!=n+1:
            st.append(tmp)
            tmp = tmp.next
            idx += 1
            
        while st:
            p.next = st.pop()
            p = p.next
        
        p.next = tmp
        return dummy.next
        

        
if __name__ == '__main__':
    head = ListNode(3)
    head.next = ListNode(5)
    su = Solution()
    su.reverseBetween(head, 1, 2)