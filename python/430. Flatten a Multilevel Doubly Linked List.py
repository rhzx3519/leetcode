"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child
"""
class Solution(object):
    def flatten(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        st = []
        pre = None
        p = head
        while p:
            pre = p
            if p.child:
                if p.next:
                    st.append(p.next)
                p.child.prev = p
                p.next= p.child
                p.child = None
                p = p.next
            else:
                p = p.next
            
            if p is None:
                if st:
                    p = st.pop()
                    pre.next = p
                    p.prev = pre
        return head

# 执行用时 :
# 32 ms
# , 在所有 Python 提交中击败了
# 92.77%
# 的用户
# 内存消耗 :
# 12.3 MB
# , 在所有 Python 提交中击败了
# 97.01%
# 的用户

