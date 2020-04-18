# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        pa = headA
        pb = headB
        while pa!=pb:
        	if pa:
        		print 'pa', pa.val 
        	if pb:
        		print 'pb', pb.val
        	pa = pa.next if pa else headB
        	pb = pb.next if pb else headA
        return pa

if __name__ == '__main__':
	t = ListNode(3)
	t.next = ListNode(4)

	headA = ListNode(1)
	headA.next = ListNode(2)
	headA.next.next = t

	headB = ListNode(11)
	headB.next = t
	
	su = Solution()
	print su.getIntersectionNode(headA, headB).val        