class Solution(object):
    def clumsy(self, N):
        """
        :type N: int
        :rtype: int
        """
        if N<=2:
            return N
        if N==3:
            return 6
        f = 1
        s = 0
        while N >= 4:
            s += (N*(N-1)/(N-2)*f + N-3)
            print N, (N*(N-1)/(N-2)*f + N-3)
            f = -1
            N -= 4
        return s - self.clumsy(N)
        
if __name__ == '__main__':
    su = Solution()
    print su.clumsy(10)