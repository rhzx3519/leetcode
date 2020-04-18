class Solution(object):
    def reorderedPowerOf2(self, N):
        """
        :type N: int
        :rtype: bool
        """
        a = [sorted(list(str(2**i))) for i in range(31)]
        return sorted(list(str(N))) in a