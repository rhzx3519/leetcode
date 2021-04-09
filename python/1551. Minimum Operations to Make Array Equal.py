class Solution(object):
    def minOperations(self, n):
        """
        :type n: int
        :rtype: int
        """
        return sum([n-i for i in range(1, n, 2)])