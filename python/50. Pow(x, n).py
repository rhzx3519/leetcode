class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        if n==0: return 1
        if x==0 or n==1: return x
        if n<0:
            return 1.0/self.myPow(x, -n)
        y = self.myPow(x, n/2)
        res = y*y
        if n&1:
            res *= x
        return res