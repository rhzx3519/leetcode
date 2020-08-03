#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class MedianFinder(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.que = []
        

    def addNum(self, num):
        """
        :type num: int
        :rtype: None
        """
        idx = self.biSearch(num)
        self.que.insert(idx, num)
        

    def findMedian(self):
        """
        :rtype: float
        """
        n = len(self.que)
        if n&1:
            return self.que[n//2]
        return (self.que[(n-1)//2] + self.que[n//2])*0.5

    def biSearch(self, val):
        l, r = 0, len(self.que)-1
        while l <= r:
            mid = l + (r-l)//2
            key = self.que[mid]
            if key==val:
                #1.先判断target是否为最后一个,是的话可以直接返回了，因为后面已经没有元素了,同时可以防止第二个条件越界
                #2.判断中点后一个和中点是否一致，不一致也可以直接返回了
            	if mid==len(self.que)-1 or key!=self.que[mid+1]:
            		return mid + 1
            	else:
            		l = mid + 1 #如果 中点后一个和中点相等，则可以从后一个开始继续二分搜索
            elif key > val:
                r = mid -1
            else:
                l = mid + 1
        return l
        

if __name__ == '__main__':
	su = MedianFinder()
	su.addNum(1)
	su.addNum(2)
	print su.findMedian()
	su.addNum(3)
	print su.findMedian()


