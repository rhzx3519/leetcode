class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

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