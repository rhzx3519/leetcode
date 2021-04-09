class Solution(object):
    def longestWPI(self, hours):
        """
        :type hours: List[int]
        :rtype: int
        """
        n = len(hours)
        arr = [1 if h>8 else -1 for h in hours]
        preSum = [0]*(n+1)
        
        for i in range(1, n+1):
            preSum[i] = arr[i-1] + preSum[i-1]
        
        st = []
        for i in range(n+1):
            if not st or (st and preSum[i] < preSum[st[-1]]):
                st.append(i)

        print arr
        print preSum
        print st

        ans = 0
        for i in range(n, -1, -1):
            while st and preSum[i] - preSum[st[-1]] > 0:
                ans = max(ans, i - st[-1])
                st.pop()
        return ans

# 那么也就是找出前缀和数组 presum 中两个索引 i 和 j，使 j - i 最大，且保证 presum[j] - presum[i] 大于 0。

        