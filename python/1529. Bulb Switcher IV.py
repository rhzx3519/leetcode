class Solution(object):
    def minFlips(self, target):
        """
        :type target: str
        :rtype: int
        """
        n = len(target)
        if n==1: return int(target[0]==1)

        n0 = n1 = 0
        for i in range(n-1, -1, -1):
            if target[i]=='0' and (i==n-1 or target[i+1]=='1'):
                n0 += 1
            elif target[i]=='1' and (i==n-1 or target[i+1]=='0'):
                n1 += 1
        return n0 + n1 - int(target[0]=='0')