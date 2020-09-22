class ProductOfNumbers(object):

    def __init__(self):
        self.nums = [1]

    def add(self, num):
        """
        :type num: int
        :rtype: None
        """
        if num==0:
            self.nums = [1]
            return
        self.nums.append(self.nums[-1] * num)



    def getProduct(self, k):
        """
        :type k: int
        :rtype: int
        """
        n = len(self.nums)
        if k >= n:
            return 0
        return self.nums[-1] / self.nums[n-k-1]




# Your ProductOfNumbers object will be instantiated and called as such:
# obj = ProductOfNumbers()
# obj.add(num)
# param_2 = obj.getProduct(k)