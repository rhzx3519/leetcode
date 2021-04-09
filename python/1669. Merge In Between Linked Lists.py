# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def mergeInBetween(self, list1, a, b, list2):
        """
        :type list1: ListNode
        :type a: int
        :type b: int
        :type list2: ListNode
        :rtype: ListNode
        """
        pa = pb = list1
        ta = 0
        tb = 0
        while ta < a - 1:
            pa = pa.next
            pb = pb.next
            ta += 1
            tb += 1

        while tb <= b:
            pb = pb.next
            tb += 1
        
        # print pa.val, pb.val
        pa.next = list2

        p2 = list2
        while p2.next:
            p2 = p2.next
        p2.next = pb

        return list1
