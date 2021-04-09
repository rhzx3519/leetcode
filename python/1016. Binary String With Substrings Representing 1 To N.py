class Solution(object):
    def queryString(self, S, N):
        """
        :type S: str
        :type N: int
        :rtype: bool
        """
        for i in range(1, N+1):
            if str(bin(i))[2:] not in S:
                return False
        return True

if __name__ == '__main__':
    S = "011010101010111101010101011111111111111111111111111111111110000000000000011111101010101001010101010101010101010101111010101010111111111111111111111111111111111100000000000000111111010101010010101010101010101010100"
    N = 1000000000
    su = Solution()
    su.queryString(S, N)