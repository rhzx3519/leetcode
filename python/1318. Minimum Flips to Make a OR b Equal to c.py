class Solution(object):
    def minFlips(self, a, b, c):
        """
        :type a: int
        :type b: int
        :type c: int
        :rtype: int
        """
        def count(n):
            c = 0
            while n:
                n = n&(n-1)
                c += 1
            return c

        d = (a|b)^c
        return count(d) + count(a&b&d)