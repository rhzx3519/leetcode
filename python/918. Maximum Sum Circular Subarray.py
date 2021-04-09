class Solution(object):
    def maxSubarraySumCircular(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        # 求但区间子数组最大和
        ans = cur = None
        for x in A:
            cur = x + max(cur, 0)
            ans = max(ans, cur)
        
        n = len(A)
        rightsum = [None]*n
        rightsum[-1] = A[-1]
        for i in range(n-2, -1, -1):
            rightsum[i] = A[i] + rightsum[i+1]
        
        maxRightsum = [None]*n
        maxRightsum[-1] = rightsum[-1]
        for i in range(n-2, -1, -1):
            maxRightsum[i] = max(rightsum[i], maxRightsum[i+1])
        # print rightsum
        # print maxRightsum
        leftSum = 0
        for i in range(n-2):
            leftSum += A[i]
            ans = max(ans, leftSum + maxRightsum[i+2])
        return ans