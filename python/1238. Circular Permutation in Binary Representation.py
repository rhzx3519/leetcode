class Solution(object):
    def circularPermutation(self, n, start):
        """
        :type n: int
        :type start: int
        :rtype: List[int]
        """
        return [start, start ^ 1] if n == 1 else self.circularPermutation(n - 1, start) + self.circularPermutation(n - 1, start ^ 1 << n - 1)[::-1] 
