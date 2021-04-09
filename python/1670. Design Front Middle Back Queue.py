class FrontMiddleBackQueue(object):

    def __init__(self):
        self.que = []


    def pushFront(self, val):
        """
        :type val: int
        :rtype: None
        """
        self.que.insert(0, val)


    def pushMiddle(self, val):
        """
        :type val: int
        :rtype: None
        """
        self.que.insert(self.midx(), val)


    def pushBack(self, val):
        """
        :type val: int
        :rtype: None
        """
        self.que.append(val)


    def popFront(self):
        """
        :rtype: int
        """
        if not self.que: return -1
        return self.que.pop(0)


    def popMiddle(self):
        """
        :rtype: int
        """
        if not self.que: return -1
        return self.que.pop(self.midx())


    def popBack(self):
        """
        :rtype: int
        """
        if not self.que: return -1
        return self.que.pop()

    def midx(self):
        return (len(self.que)+1)/2


if __name__ == '__main__':
    su = FrontMiddleBackQueue()
    su.pushFront(1)
    su.pushBack(2)
    su.pushMiddle(3)
    su.pushMiddle(4)
    su.popMiddle()
    su.popMiddle()
    su.popBack()
    su.popFront()
    print su.que

# Your FrontMiddleBackQueue object will be instantiated and called as such:
# obj = FrontMiddleBackQueue()
# obj.pushFront(val)
# obj.pushMiddle(val)
# obj.pushBack(val)
# param_4 = obj.popFront()
# param_5 = obj.popMiddle()
# param_6 = obj.popBack()