# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        dummy = ListNode(0)
        p = dummy
        while True:
            count = k
            stack = []
            tmp = head
            while count and tmp:
                count -= 1
                stack.append(tmp)
                tmp = tmp.next
            if count:
                p.next = head
                break
            while stack: # 反转K个结点
                p.next = stack.pop()
                p = p.next
            p.next = tmp # 连接剩下的链表
            head = tmp
        return dummy.next


class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        dummy = ListNode(0)
        dummy.next = head
        p = dummy
        k = 2
        st = []
        while head or st:
            if k > 0 and head:
                k -= 1
                st.append(head)
                head = head.next
                continue
            tmp = p
            # print st, head
            while st:
                node = st.pop()
                tmp.next = node
                tmp = tmp.next 
            p = tmp   
            p.next = head
            # print p, p.next        
            k = 2
        # print st
        return dummy.next


        
l1 = ListNode(1)
l1.next = ListNode(2)
l1.next.next = ListNode(3)
l1.next.next.next = ListNode(4)
l1.next.next.next.next = ListNode(5)
s = Solution()
p = s.reverseKGroup(l1, 1)
while p:
    print p.val

        