class StockSpanner(object):

    def __init__(self):
        self.stock = []
        self.span = []

    def next(self, price):
        """
        :type price: int
        :rtype: int
        """
        w = 1
        while self.stock and self.stock[-1] <= price:
            self.stock.pop()
            w += self.span.pop()
        self.stock.append(price)
        self.span.append(w)
        # print self.span, self.stock
        return self.span[-1]




# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)