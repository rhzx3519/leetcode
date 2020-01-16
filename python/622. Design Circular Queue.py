class MyCircularQueue(object):
    
    def __init__(self, k):
        """
        Initialize your data structure here. Set the size of the queue to be k.
        :type k: int
        """
        self.a = [0]*(k+1)
        self.k = k+1
        self.start = 0
        self.end = 0
        

    def enQueue(self, value):
        """
        Insert an element into the circular queue. Return true if the operation is successful.
        :type value: int
        :rtype: bool
        """
        if self.isFull():
        	return False
        self.a[self.end] = value
        self.end = (self.end+1)%self.k
        return True
        

    def deQueue(self):
        """
        Delete an element from the circular queue. Return true if the operation is successful.
        :rtype: bool
        """
        if self.isEmpty():
        	return False
        self.start = (self.start+1)%self.k
        return True
        

    def Front(self):
        """
        Get the front item from the queue.
        :rtype: int
        """
        if self.isEmpty():
        	return -1
        return self.a[self.start]
        

    def Rear(self):
        """
        Get the last item from the queue.
        :rtype: int
        """
        if self.isEmpty():
        	return -1
        return self.a[(self.end-1)%self.k]


    def isEmpty(self):
        """
        Checks whether the circular queue is empty or not.
        :rtype: bool
        """
        return self.start==self.end
        

    def isFull(self):
        """
        Checks whether the circular queue is full or not.
        :rtype: bool
        """
        return (self.end + 1)%self.k == self.start


# Your MyCircularQueue object will be instantiated and called as such:
k = 3
obj = MyCircularQueue(k)
param_1 = obj.enQueue(1)
obj.enQueue(2)
obj.enQueue(3)
print obj.a, obj.start, obj.end
# param_2 = obj.deQueue()
print obj.a, obj.start, obj.end
param_3 = obj.Front()
param_4 = obj.Rear()
print param_3, param_4

# obj.deQueue()
# obj.deQueue()
param_5 = obj.isEmpty()
param_6 = obj.isFull()
print param_5, param_6
