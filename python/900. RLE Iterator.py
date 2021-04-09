class RLEIterator(object):

    def __init__(self, A):
        """
        :type A: List[int]
        """
        self.A = A
        

    def next(self, n):
        """
        :type n: int
        :rtype: int
        """
        if not self.A:
            return -1
        digit = self.A[1]
        while n > 0 and self.A:
            num, digit = self.A[0], self.A[1]
            n -= num
            if n >= 0:
                self.A.pop(0)
                digit = self.A.pop(0)
            else:
                self.A[0] = -n
        # print n, self.A
        return -1 if n > 0 else digit



        


# Your RLEIterator object will be instantiated and called as such:
# obj = RLEIterator(A)
# param_1 = obj.next(n)