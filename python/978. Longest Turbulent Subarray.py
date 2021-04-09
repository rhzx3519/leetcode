class Solution(object):
    def maxTurbulenceSize(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        n = len(arr)
        if n <= 1: return n
        ans = 0
        s = 0

        last = 0
        for i in xrange(1, n):
            if arr[i-1] > arr[i]:
                cur = 1
            elif arr[i-1] < arr[i]:
                cur = -1
            else:
                cur = 0
            
            if last < 0 and cur > 0:
                s += 1
            elif last > 0 and cur < 0:
                s += 1
            elif cur != 0:
                s = 1
            else:
                s = 0
            # print i, s
            ans = max(ans, s+1)
            last = cur
        return ans
            
