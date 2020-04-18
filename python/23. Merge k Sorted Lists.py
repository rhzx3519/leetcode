# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

import heapq

class PriorityQueue(object):
    def __init__(self):
        self._queue = []
        self._index = 0
        
    def push(self, item, priority):
        heapq.heappush(self._queue, (priority, self._index, item))
        self._index += 1
        
    def pop(self):
        return heapq.heappop(self._queue)[-1]
    
    def size(self):
        return len(self._queue)

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        if not lists:
            return None
        st = PriorityQueue()
        for x in lists:
            if x:
                st.push(x, x.val)
        if st.size() == 0:
            return None
            
        head = tail = None
        while st.size() != 0:
            ptr = st.pop()
            if ptr.next:
                st.push(ptr.next, ptr.next.val)
            ptr.next = None
            
            if ptr == None:
                continue
            if tail == None:
                head = ptr
            else:
                tail.next = ptr
            tail = ptr
            tail.next = None
        return head


# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from Queue import PriorityQueue

class Comparable(object):

    def __init__(self, node):
        self.node = node
    
    def __cmp__(self, other):
        if self.node.val < other.node.val:
            return -1
        elif self.node.val == other.node.val:
            return 0
        else:
            return 1

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        que = PriorityQueue()
        for node in [i for i in lists if i]:
            que.put(Comparable(node))

        head = tail = None
        while que.qsize()>0:
            p = que.get().node
            if head is None:
                head = p
                tail = p
            else:
                tail.next = p
                tail = p
            if p.next:
                que.put(Comparable(p.next))

        return head












    
        