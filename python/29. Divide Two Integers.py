class Solution(object):
    def divide(self, dividend, divisor):
        """
        :type dividend: int
        :type divisor: int
        :rtype: int
        """
        sign = 1
        n1, n2 = dividend, divisor
        if n1 == 0 or abs(n1) < abs(n2):
            return 0
        if (n1 > 0 and n2 < 0) or (n1 < 0 and n2 > 0):
            sign = 0
        n1, n2 = abs(n1), abs(n2)
        res = 0
        while n1 >= n2:
            t = 1
            b = n2
            while n1 > b<<1:
                b = b<<1
                t = t<<1
            res += t
            n1 -= b
        res = sign and res or -res
        if res > pow(2,31)-1 or res < -pow(2,31):
            return pow(2,31)
        return res
            
            
s = Solution()
print s.divide(3, -2)        