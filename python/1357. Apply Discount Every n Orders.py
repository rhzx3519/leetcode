class Cashier(object):

    def __init__(self, n, discount, products, prices):
        """
        :type n: int
        :type discount: int
        :type products: List[int]
        :type prices: List[int]
        """
        self.n = n
        self.discount = 1 - discount*0.01
        self.d = {}
        for i, j in zip(products, prices):
            self.d[i] = j
        self.cur = 0


    def getBill(self, product, amount):
        """
        :type product: List[int]
        :type amount: List[int]
        :rtype: float
        """
        self.cur += 1
        discount = self.discount if not self.cur%self.n else 1
        ans = 0
        for i, j in zip(product, amount):
            ans += self.d[i]*j

        
        return ans * discount
        



# Your Cashier object will be instantiated and called as such:
# obj = Cashier(n, discount, products, prices)
# param_1 = obj.getBill(product,amount)