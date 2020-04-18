class Solution(object):
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        a = [1]
        for _ in range(0, rowIndex):
            b = [1]
            for i in range(1, len(a)):
                b.append(a[i-1]+a[i])
            b.append(1)
            a = b
        return a

if __name__ == '__main__':
    su = Solution()
    print su.getRow(3)