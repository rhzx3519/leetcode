class Solution(object):
    def minKBitFlips(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        que = []
        N = len(A)
        ans = 0
        for i in range(N):
            if que and i >= que[0] + K:
                que.pop(0)
            if len(que)%2 == A[i]:
                if i + K > N: return -1
                que.append(i)
                ans += 1
        return ans

# https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/solution/hua-dong-chuang-kou-shi-ben-ti-zui-rong-z403l/
