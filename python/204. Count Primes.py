class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """
        a = [1]*(n+1)
        count = 0
        for i in range(2, n):
            if a[i]:
                count += 1
            j = 2
            while i*j<n:
                a[i*j] = 0
                j += 1

        return count