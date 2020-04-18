# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def nextLargerNodes(self, head):
        """
        :type head: ListNode
        :rtype: List[int]
        """
        p = head
        res = []
        st = []
        i = 0
        while p:
            while st and p.val > res[st[-1]]:
                res[st[-1]] = p.val
                st.pop()
            res.append(p.val)
            st.append(i)
            i += 1
            p = p.next

        while st:
            res[st[-1]] = 0
            st.pop()
            
        return res

if __name__ == '__main__':
    head = ListNode(2)
    head.next = ListNode(1)
    head.next.next = ListNode(5)

    su = Solution()
    print su.nextLargerNodes(head)          